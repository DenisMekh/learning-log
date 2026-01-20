package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	WriteJSON(http.StatusOK, []byte(`{"message": "pong"}`), w)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	WriteJSON(http.StatusOK, []byte(`{"message": "ok"}`), w)
}

func (h *Handler) SomeLong(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	var number int
	for i := 0; i < 10_000_000; i++ {
		number++
	}
	select {
	case <-time.After(time.Second):
		WriteJSON(http.StatusOK, []byte(`{"message": "ok", "number": "someNumber"}`), w)
	case <-ctx.Done():
		WriteJSON(http.StatusInternalServerError, []byte(`{"message": "context deadline exceeded"}`), w)
	}
}

func WriteJSON(status int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}
}
