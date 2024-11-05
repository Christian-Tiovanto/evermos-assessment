package errors

// ForbiddenGeneralCode defines forbidden error general default code
const ForbiddenGeneralCode = "forbidden_error"

// Forbidden creates forbidden error
func Forbidden(msg string) *Error {
	return &Error{tipe: ForbiddenError, message: msg, code: ForbiddenGeneralCode}
}
