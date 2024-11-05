package errors

// AuthenticationGeneralCode defines authentication error general default code
const AuthenticationGeneralCode = "authentication_error"

// Authentication creates authentication error
func Authentication(msg string) *Error {
	return &Error{tipe: AuthenticationError, message: msg, code: AuthenticationGeneralCode}
}
