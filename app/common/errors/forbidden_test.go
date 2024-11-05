package errors_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/errors"
)

func TestForbidden(t *testing.T) {
	err := errors.Forbidden(defaultMessage)
	assert.Equal(t, defaultMessage, err.Message())
	assert.Equal(t, errors.ForbiddenError, err.Type())
	assert.Equal(t, errors.ForbiddenGeneralCode, err.Code())
	assert.Equal(t, http.StatusForbidden, err.HTTPStatus())
}
