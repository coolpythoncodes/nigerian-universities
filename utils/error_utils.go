package utils

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   true,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   true,
	}
}
