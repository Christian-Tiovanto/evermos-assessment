// Package errors provides service error definition.
package errors

import (
	"bytes"
	stderrors "errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

// ErrorType defines error type
type ErrorType int

const (
	// Unidentified indicates the error type is unidentified
	Unidentified ErrorType = iota
	// InternalServerError indicates internal server error
	InternalServerError
	// BadRequestError indicates bad request error
	BadRequestError
	// AuthenticationError indicates authentication error
	AuthenticationError
	// ForbiddenError indicates forbidden error
	ForbiddenError
	// NotFoundError indicates not found error
	NotFoundError
	// CancelledError indicates cancelled error
	CancelledError
)

// Error defines service error
type Error struct {
	originErr error
	code      string
	message   string
	tipe      ErrorType
}

const stacktraceMaxLength = 1000

// Error returns error string to satisfy error interface
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString(e.message)
	if e.code != "" {
		buf.WriteString(fmt.Sprintf(" (code: %s)", e.code))
	}
	if e.originErr != nil {
		st := fmt.Sprintf(" (origin err stacktrace: %+v", e.originErr)

		buf.WriteString(fmt.Sprintf(" (origin err message: %s)", e.originErr))
		buf.WriteString(st[:min(len(st), stacktraceMaxLength)])
		if stacktraceMaxLength < len(st) {
			buf.WriteString("... SNIPPED ....)")
		}
	}
	return buf.String()
}

// Type returns error type
func (e *Error) Type() ErrorType {
	if e == nil {
		return Unidentified
	}
	return e.tipe
}

// Message returns error message for client
func (e *Error) Message() string {
	if e == nil {
		return ""
	}
	return e.message
}

// Code returns error code
func (e *Error) Code() string {
	if e == nil {
		return ""
	}
	return e.code
}

// String returns error string
func (e *Error) String() string {
	return e.Error()
}

// WithErr sets origin error
func (e *Error) WithErr(err error) *Error {
	if e == nil {
		return nil
	}
	e.originErr = err
	return e
}

// WithCode sets error code
func (e *Error) WithCode(code string) *Error {
	if e == nil {
		return nil
	}
	e.code = code
	return e
}

// Unwrap returns origin error
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.originErr
}

// StatusClientClosedRequest is HTTP status code for client closed request case.
const StatusClientClosedRequest = 499

var errorTypeToHTTPStatus = map[ErrorType]int{
	InternalServerError: http.StatusInternalServerError,
	BadRequestError:     http.StatusBadRequest,
	AuthenticationError: http.StatusUnauthorized,
	ForbiddenError:      http.StatusForbidden,
	NotFoundError:       http.StatusNotFound,
	CancelledError:      StatusClientClosedRequest,
}

// HTTPStatus returns HTTP status code from error type
func (e *Error) HTTPStatus() int {
	if e == nil {
		return http.StatusOK
	}
	status, registered := errorTypeToHTTPStatus[e.tipe]
	if !registered {
		return http.StatusInternalServerError
	}
	return status
}

var errorTypeToGrpcCode = map[ErrorType]codes.Code{
	InternalServerError: codes.Internal,
	BadRequestError:     codes.InvalidArgument,
	AuthenticationError: codes.Unauthenticated,
	ForbiddenError:      codes.PermissionDenied,
	NotFoundError:       codes.NotFound,
	CancelledError:      codes.Canceled,
}

// GRPCCode returns gRPC status code from error type
func (e *Error) GRPCCode() codes.Code {
	if e == nil {
		return codes.OK
	}
	c, registered := errorTypeToGrpcCode[e.tipe]
	if !registered {
		return codes.Unknown
	}
	return c
}

// Unwrap call Go standard errors package Unwrap function.
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}

// Is call Go standard errors package Is function.
func Is(err, target error) bool {
	return stderrors.Is(err, target)
}

// As call Go standard errors package As function.
func As(err error, target any) bool {
	return stderrors.As(err, target)
}
