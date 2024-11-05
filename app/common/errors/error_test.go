package errors_test

import (
	"net/http"
	"testing"

	goredis "github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/errors"
)

const defaultMessage = "my message"

func TestError(t *testing.T) {
	originErr := pkgerrors.New("error string")

	t.Run("when error nil", func(t *testing.T) {
		t.Parallel()
		var err *errors.Error

		assert.Empty(t, err.Error())
		assert.Empty(t, err.String())
		assert.Empty(t, err.Code())
		assert.Empty(t, err.Message())

		assert.Nil(t, err.WithErr(originErr))
		assert.Nil(t, err.WithCode("code"))
		assert.Nil(t, err.Unwrap())

		assert.Equal(t, errors.Unidentified, err.Type())
		assert.Equal(t, http.StatusOK, err.HTTPStatus())
	})

	t.Run("standard error func", func(t *testing.T) {
		t.Parallel()

		err := errors.BadRequest(defaultMessage)

		assert.Contains(t, err.Error(), defaultMessage)
		assert.Contains(t, err.Error(), errors.BadRequestGeneralCode)
		assert.NotContains(t, err.Error(), originErr.Error())
		assert.Nil(t, err.Unwrap())
		assert.Equal(t, err.Error(), err.String())
	})

	t.Run("error func with custom code", func(t *testing.T) {
		t.Parallel()

		customCode := "custom_error_code"
		err := errors.BadRequest(defaultMessage).WithCode(customCode)

		assert.Contains(t, err.Error(), defaultMessage)
		assert.Contains(t, err.Error(), customCode)
		assert.NotContains(t, err.Error(), originErr.Error())
		assert.Nil(t, err.Unwrap())
		assert.Equal(t, err.Error(), err.String())
	})

	t.Run("error func with origin error", func(t *testing.T) {
		t.Parallel()

		err := errors.BadRequest(defaultMessage).WithErr(originErr)

		assert.Contains(t, err.Error(), defaultMessage)
		assert.Contains(t, err.Error(), errors.BadRequestGeneralCode)
		assert.Contains(t, err.Error(), originErr.Error())
		assert.Equal(t, err.Error(), err.String())

		// error stack trace should be included in error message
		assert.Contains(t, err.Error(), "app/common/errors/error_test.go")

		if assert.NotNil(t, err.Unwrap()) {
			assert.True(t, pkgerrors.Is(err, originErr))
		}
	})

	t.Run("http status unregistered type", func(t *testing.T) {
		t.Parallel()

		err := errors.Error{}
		assert.Equal(t, http.StatusInternalServerError, err.HTTPStatus())
	})

	t.Run("standard error functions", func(t *testing.T) {
		t.Parallel()

		rediserr := goredis.Nil
		pkgerr := pkgerrors.WithStack(rediserr)
		svcerr := errors.InternalServer(pkgerr)

		assert.True(t, errors.Is(pkgerr, rediserr))
		assert.True(t, errors.Is(svcerr, rediserr))

		var err *errors.Error
		fn := func() error {
			return svcerr
		}
		assert.True(t, errors.As(fn(), &err))
		assert.Equal(t, errors.InternalServerError, err.Type())

		stderr := errors.Unwrap(svcerr)
		assert.Equal(t, pkgerr, stderr)
		assert.Equal(t, rediserr, errors.Unwrap(stderr))
	})
}
