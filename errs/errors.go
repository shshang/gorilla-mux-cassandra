package errs

import "net/http"

// The code is taken care of by the HTTP error code.
// In the body of the HTTP error, we only need "message", not "code"
// Remove "Code" from the JSON output of AppError
type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
