package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/mughieams/evermos-assessment/app/common/config"
)

// Subscription is an interface to interact with rabbitMQ
type Subscription interface {
	QueueDeclare(queue string) error
	Publish(queue string, message []byte) error
	Consume(queue string, consumerFunc func([]byte)) error
	Close() error
}

type rabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

var _ Subscription = (*rabbitMQ)(nil)

// NewSubscription is a constructor to create new rabbitMQ subscription
func NewSubscription(cfg *config.RabbitMQ) (Subscription, error) {
	conn, err := amqp.Dial(cfg.ConnectionString())
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitMQ{conn: conn, channel: ch}, nil
}

func (rb *rabbitMQ) QueueDeclare(queue string) error {
	_, err := rb.channel.QueueDeclare(queue, true, false, false, false, nil)
	return err
}

// Publish is a function to publish message to rabbitMQ
func (rb *rabbitMQ) Publish(queue string, message []byte) error {
	_, err := rb.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = rb.channel.Publish(
		"",    // exchange
		queue, // key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	return err
}

// Consume is a function to consume message from rabbitMQ
func (rb *rabbitMQ) Consume(queue string, consumerFunc func([]byte)) error {
	msgs, err := rb.channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			consumerFunc(msg.Body)
		}
	}()

	return nil
}

// Close is a function to close rabbitMQ connection
func (rb *rabbitMQ) Close() error {
	err := rb.channel.Close()
	if err != nil {
		return err
	}

	err = rb.conn.Close()
	if err != nil {
		return err
	}

	return nil
}
