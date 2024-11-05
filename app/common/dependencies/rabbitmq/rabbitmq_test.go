package rabbitmq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/config"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/rabbitmq"
)

func TestNewRabbitMQ(t *testing.T) {
	cfg := &config.RabbitMQ{
		Host:     "localhost",
		Port:     "5672",
		User:     "user",
		Password: "password",
	}

	t.Run("fail build rabbitmq client", func(t *testing.T) {
		client, err := rabbitmq.NewSubscription(cfg)

		assert.NotNil(t, err)
		assert.Nil(t, client)
	})
}
