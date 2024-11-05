package errors

// BadRequestGeneralCode defines bad request error general default code
const BadRequestGeneralCode = "bad_request_error"

// BadRequest creates bad request error
func BadRequest(msg string) *Error {
	return &Error{tipe: BadRequestError, message: msg, code: BadRequestGeneralCode}
}
