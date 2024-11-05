package errors_test

import (
	"net/http"
	"testing"

	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/mughieams/evermos-assessment/app/common/errors"
)

func TestInternalServer(t *testing.T) {
	err := errors.InternalServer(pkgerrors.New("error string"))

	assert.Equal(t, errors.InternalServerMessage, err.Message())
	assert.Equal(t, errors.InternalServerError, err.Type())
	assert.Equal(t, errors.InternalServerGeneralCode, err.Code())
	assert.Equal(t, http.StatusInternalServerError, err.HTTPStatus())

	// error stack trace should be included in error message
	assert.Contains(t, err.Error(), "app/common/errors/internal_server_test.go")
}
