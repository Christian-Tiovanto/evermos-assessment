package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/config"
)

func TestNewConfig(t *testing.T) {
	t.Run("successfully load config", func(t *testing.T) {
		cfg, err := config.Load("../../../env.sample")
		assert.Nil(t, err)
		assert.NotNil(t, cfg)
	})
}
