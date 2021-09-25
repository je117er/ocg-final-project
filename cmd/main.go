package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/je117er/ocg-final-project/internal/rabbitmq"
	"github.com/je117er/ocg-final-project/internal/scheduler"
	"github.com/je117er/ocg-final-project/internal/server"
	"github.com/je117er/ocg-final-project/internal/utils"
	"log"
)

var logger = utils.SugarLog()

func init() {
	logger.Info("Init .....")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"admin", "123456Aa@", "localhost", "6033", "vaccine-covid-19"))
	if err != nil {
		logger.Fatal(err)
	}

	ctx := context.Background()
	rmqURI := fmt.Sprintf("amqp://%s:%s@%s:%d", "admin", "123456Aa@", "localhost", 5672)
	exch := "order"
	exchType := "direct"
	queue := "order_processor"
	routingKey := ""

	// connect to rabbitmq
	rmq := rabbitmq.NewRMQ(rmqURI)

	// create 1 channel for producer
	pCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Can't get channel: ", err)
		return
	}
	// create 1 channel for consumer
	cCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Can't get channel: ", err)
		return
	}

	producer := rabbitmq.NewProducer(ctx, pCh, exch, exchType, routingKey)
	consumer := rabbitmq.NewConsumer(ctx, cCh, exch, exchType, routingKey, queue, db)

	scheduler := scheduler.NewScheduler(ctx, db, producer)
	go scheduler.Start()
	go consumer.Start()
}

func main() {
	logger.Info("Start application...")

	log.Fatal(server.InitServer())
}
