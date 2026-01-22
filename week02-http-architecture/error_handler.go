package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

type ErrorJSON struct {
	Error     string `json:"error"`
	RequestID string `json:"request_id"`
}

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		HandlerError(w, r, err)
	}

}

func HandlerError(w http.ResponseWriter, r *http.Request, err error) {
	requestID := GetRequestID(r.Context())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var statusCode int
	var ae *AppError
	if errors.As(err, &ae) {
		statusCode = ae.StatusCode
	} else if errors.Is(err, ErrNotFound) {
		statusCode = http.StatusNotFound
	} else if errors.Is(err, ErrBadRequest) {
		statusCode = http.StatusBadRequest
	} else {
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(ErrorJSON{Error: err.Error(), RequestID: requestID})
}
