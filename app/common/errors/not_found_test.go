package errors_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/errors"
)

func TestNotFound(t *testing.T) {
	err := errors.NotFound(defaultMessage)
	assert.Equal(t, defaultMessage, err.Message())
	assert.Equal(t, errors.NotFoundError, err.Type())
	assert.Equal(t, errors.NotFoundGeneralCode, err.Code())
	assert.Equal(t, http.StatusNotFound, err.HTTPStatus())
}
