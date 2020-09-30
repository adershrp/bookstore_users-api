package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"errors"`
}

// Bad request error
func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   http.StatusText(http.StatusBadRequest),
	}
}

// Not found error
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   http.StatusText(http.StatusBadRequest),
	}
}

// Internal Server Error
func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   http.StatusText(http.StatusInternalServerError),
	}
}
