package errors_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/errors"
)

func TestBadRequest(t *testing.T) {
	err := errors.BadRequest(defaultMessage)
	assert.Equal(t, defaultMessage, err.Message())
	assert.Equal(t, errors.BadRequestError, err.Type())
	assert.Equal(t, errors.BadRequestGeneralCode, err.Code())
	assert.Equal(t, http.StatusBadRequest, err.HTTPStatus())
}
