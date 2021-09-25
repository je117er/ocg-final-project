package rabbitmq

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	bookingRepository "github.com/je117er/ocg-final-project/internal/booking/repository"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/sendmail"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/streadway/amqp"
)

type Consumer struct {
	ctx        context.Context
	channel    *amqp.Channel
	queue      string
	exchange   string
	exchType   string
	bindingKey string
	db         *sql.DB
}

var logger = utils.SugarLog()

func NewConsumer(ctx context.Context, channel *amqp.Channel, exchange, exchType, bindingKey, queue string, db *sql.DB) *Consumer {
	return &Consumer{
		ctx:        ctx,
		channel:    channel,
		exchange:   exchange,
		exchType:   exchType,
		bindingKey: bindingKey,
		queue:      queue,
		db:         db,
	}
}

func (consumer *Consumer) declare() error {
	// declare exchange
	if err := consumer.channel.ExchangeDeclare(
		consumer.exchange,
		consumer.exchType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	// declare queue
	queue, err := consumer.channel.QueueDeclare(
		consumer.queue,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil
	}

	if err = consumer.channel.QueueBind(
		queue.Name, consumer.bindingKey, consumer.exchange, false, nil); err != nil {
		return err
	}
	return nil
}

func (consumer *Consumer) Start() {
	logger.Info("Start consumer...")
	consumer.declare()

	delivery, err := consumer.channel.Consume(
		consumer.queue,
		"",
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		return
	}

	for {
		select {
		case mess := <-delivery:
			mess.Ack(false)
			var sendMail models.SentMail
			_ = json.Unmarshal(mess.Body, &sendMail)
			_, err := sendmail.SendEmailThankYou(sendMail.Email, sendMail.Email)
			if err != nil {
				fmt.Println("error: ", err)
				continue
			}

			bookingRepo := bookingRepository.NewBookingRepository(consumer.db)
			bookingRepo.UpdateIsSendRemindEmail(consumer.ctx, sendMail.BookingID)

		}
	}
}

func (consumer *Consumer) Close() error {
	return consumer.channel.Close()
}
