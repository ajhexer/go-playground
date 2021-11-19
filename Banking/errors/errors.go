package errors

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func ValidationError(message string) *AppError{
	return &AppError{
		http.StatusBadRequest,
		message,
	}
}

