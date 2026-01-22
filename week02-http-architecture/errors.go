package main

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

var (
	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal server error")
)

func (e *AppError) Error() string { return e.Message }

func (e *AppError) Unwrap() error { return e.Err }

func NewBadRequest(err error, msg string) error {
	return &AppError{
		Err:        err,
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NewNotFound(err error, msg string) error {
	return &AppError{
		Err:        err,
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}
func Wrapf(err error, format string, args ...interface{}) error {
	return fmt.Errorf("%s %w", fmt.Sprintf(format, args...), err)
}
