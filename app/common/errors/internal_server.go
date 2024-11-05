package errors

import (
	"context"
	stderrors "errors"
)

const (
	// InternalServerMessage defines internal server error message
	InternalServerMessage = "Internal server error"
	// CancelledMessage defines error message when context cancelled
	CancelledMessage = "Request cancelled"

	// InternalServerGeneralCode defines internal server error general default code
	InternalServerGeneralCode = "internal_server_error"
	// CancelledCode defines internal server error code when context cancelled
	CancelledCode = "context_canceled"
)

// InternalServer creates internal server error
func InternalServer(err error) *Error {
	if stderrors.Is(err, context.Canceled) {
		return &Error{tipe: CancelledError, message: CancelledMessage, code: CancelledCode, originErr: err}
	}
	return &Error{tipe: InternalServerError, message: InternalServerMessage, code: InternalServerGeneralCode, originErr: err}
}
