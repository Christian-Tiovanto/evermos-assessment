package errors_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/errors"
)

func TestAuthentication(t *testing.T) {
	err := errors.Authentication(defaultMessage)
	assert.Equal(t, defaultMessage, err.Message())
	assert.Equal(t, errors.AuthenticationError, err.Type())
	assert.Equal(t, errors.AuthenticationGeneralCode, err.Code())
	assert.Equal(t, http.StatusUnauthorized, err.HTTPStatus())
}
