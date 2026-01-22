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
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			WriteJSON(http.StatusOK, []byte(`{"message": "ok", "number": "someNumber"}`), w)

		default:
			number++
		}
	}

}

func (h *Handler) SomeError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, ErrNotFound.Error(), http.StatusNotFound)
}

func WriteJSON(status int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}

}
