package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	value := r.Header.Get("X-Request-ID")
	now := time.Now().UTC().Format(time.RFC3339)
	var full map[string]any
	if value == "" {
		full = map[string]any{"time": now}
	} else {
		full = map[string]any{"time": now, "value": value}
	}
	if err := json.NewEncoder(w).Encode(full); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	status := map[string]interface{}{"status": "ok"}
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func LongPooling(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	select {
	case <-time.After(5 * time.Second):
		w.WriteHeader(http.StatusServiceUnavailable)
	case <-ctx.Done():
		return
	}
}
