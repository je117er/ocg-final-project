package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/streadway/amqp"
)

type Producer struct {
	ctx        context.Context
	channel    *amqp.Channel
	exchange   string
	exchType   string
	routingKey string
	dataChan   <-chan *models.SentMail
}

func NewProducer(ctx context.Context, channel *amqp.Channel, exchange, exchType, routingKey string, dataChan <-chan *models.SentMail) *Producer {
	return &Producer{
		ctx:        ctx,
		channel:    channel,
		exchange:   exchange,
		exchType:   exchType,
		routingKey: routingKey,
		dataChan:   dataChan,
	}
}

func (producer *Producer) Start() {
	logger.Info("Start producer...")
	producer.declare()

	for {
		data, _ := <-producer.dataChan
		producer.Public(data)
	}
}

func (producer *Producer) Public(sendMail *models.SentMail) error {
	jsonData, _ := json.Marshal(sendMail)

	if err := producer.channel.Publish(
		producer.exchange,
		producer.routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            jsonData,
			DeliveryMode:    amqp.Transient,
		},
	); err != nil {
		return err
	}

	return nil
}

func (producer *Producer) declare() error {
	if err := producer.channel.ExchangeDeclare(
		producer.exchange,
		producer.exchType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}
	return nil
}
