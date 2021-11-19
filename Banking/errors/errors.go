package errors

import "net/http"

type AppError struct {
	code 	int
	message string
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		code: http.StatusInternalServerError,
		message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		code: http.StatusNotFound,
		message: message,
	}
}

func ValidationError(message string) *AppError{
	return &AppError{
		http.StatusBadRequest,
		message,
	}
}

