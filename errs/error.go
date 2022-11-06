package errs

import "net/http"

type AppError struct {
	Code int    `json:",omitempty"`
	Msg  string `json:"message,omitempty"`
}

func NewRowNotFoundError(msg string) AppError {
	return AppError{Code: 404, Msg: msg}
}

func NewDataBaseError(msg string) AppError {
	return AppError{Code: 500, Msg: msg}
}

func NewValidationError(msg string) AppError {
	return AppError{Code: http.StatusUnprocessableEntity, Msg: msg}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Msg: e.Msg,
	}
}
