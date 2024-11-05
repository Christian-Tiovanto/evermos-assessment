package postgresql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/config"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
)

func TestNewPool(t *testing.T) {
	cfg := &config.PostgreSQL{
		Host:            "localhost",
		Port:            "5432",
		DBName:          "postgres",
		User:            "user",
		Password:        "password",
		MaxOpenConns:    "10",
		MaxConnLifetime: "10m",
		MaxIdleLifetime: "5m",
	}

	t.Run("fail build sql client", func(t *testing.T) {
		client, err := postgresql.NewPool(cfg)

		assert.NotNil(t, err)
		assert.Nil(t, client)
	})
}
