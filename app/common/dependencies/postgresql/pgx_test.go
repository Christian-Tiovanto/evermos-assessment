package postgresql_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
)

func Test_PgErrCode(t *testing.T) {
	t.Run("when pgconn", func(t *testing.T) {
		err := &pgconn.PgError{Code: postgresql.ErrFKViolationCode}
		assert.Equal(t, postgresql.ErrFKViolationCode, postgresql.PgErrCode(err))
	})

	t.Run("another error", func(t *testing.T) {
		assert.Empty(t, postgresql.PgErrCode(errors.New("internal server error")))
	})
}
