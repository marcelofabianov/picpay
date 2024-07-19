package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/picpay/config"
	"github.com/marcelofabianov/picpay/pkg/rabbitmq"
	"github.com/marcelofabianov/picpay/pkg/zap"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	logger, err := zap.NewLogger(cfg.Log)
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}
	defer logger.Close()

	rmq, err := rabbitmq.NewRabbitMQ(cfg.MessageBroker.Url)
	if err != nil {
		logger.Fatal("Failed to create RabbitMQ client")
	}
	defer rmq.Close()

	exchangeConfig := rabbitmq.ExchangeQueueConfig{
		ExchangeName:       "transfer_exchange",
		ExchangeType:       "direct",
		ExchangeDurable:    true,
		ExchangeAutoDelete: false,
		ExchangeInternal:   false,
		ExchangeNoWait:     false,
		ExchangeArgs:       nil,
	}

	err = rmq.DeclareExchange(exchangeConfig)
	if err != nil {
		logger.Fatal("Failed to declare exchange")
	}

	pubConfig := rabbitmq.PublishConfig{
		Exchange:  "transfer_exchange",
		Key:       "email",
		Mandatory: false,
		Immediate: false,
	}

	err = rmq.Publish(context.Background(), pubConfig, []byte("Email notification"))
	if err != nil {
		logger.Fatal("Failed to publish message")
	}

	consConfig := rabbitmq.ConsumeConfig{
		Exchange:          "transfer_exchange",
		Key:               "email",
		QueueName:         "email_notifications_queue",
		Durable:           true,
		Exclusive:         false,
		AutoDelete:        false,
		Args:              nil,
		NoAck:             false,
		ExclusiveConsumer: false,
		NoLocal:           false,
		Wait:              false,
	}

	err = rmq.Consumer(context.Background(), consConfig, func(message []byte) error {
		logger.Info("Received message")
		logger.Info(string(message))
		return nil
	})
	if err != nil {
		logger.Fatal("Failed to consume message")
	}

	logger.Info("Application started!")
}
