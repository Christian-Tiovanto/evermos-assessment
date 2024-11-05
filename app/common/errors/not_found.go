package errors

// NotFoundGeneralCode defines not found error general default code
const NotFoundGeneralCode = "not_found"

// NotFound creates not found error
func NotFound(msg string) *Error {
	return &Error{tipe: NotFoundError, message: msg, code: NotFoundGeneralCode}
}
