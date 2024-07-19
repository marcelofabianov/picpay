package rabbitmq

import (
	"context"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

type PublishConfig struct {
	Exchange  string
	Key       string
	Mandatory bool
	Immediate bool
}

type ConsumeConfig struct {
	Exchange          string
	Key               string
	QueueName         string
	Durable           bool
	Exclusive         bool
	AutoDelete        bool
	Args              amqp.Table
	NoAck             bool
	ExclusiveConsumer bool
	NoLocal           bool
	Wait              bool
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{conn: conn, ch: ch}, nil
}

func (r *RabbitMQ) Publish(ctx context.Context, config PublishConfig, message []byte) error {
	err := r.ch.Publish(
		config.Exchange,
		config.Key,
		config.Mandatory,
		config.Immediate,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQ) Consumer(ctx context.Context, config ConsumeConfig, handler func(message []byte) error) error {
	q, err := r.ch.QueueDeclare(
		config.QueueName,
		config.Durable,
		config.Exclusive,
		config.AutoDelete,
		config.Wait,
		config.Args,
	)
	if err != nil {
		return err
	}

	err = r.ch.QueueBind(
		q.Name,
		config.Key,
		config.Exchange,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := r.ch.Consume(
		q.Name,
		"",
		config.NoAck,
		config.ExclusiveConsumer,
		config.NoLocal,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		err := handler(msg.Body)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RabbitMQ) Close() error {
	if err := r.ch.Close(); err != nil {
		return err
	}
	return r.conn.Close()
}
