package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	start := time.Now()
	value := r.Header.Get("X-Request-ID")
	var full map[string]any
	if value == "" {
		full = map[string]any{"time": start}
	} else {
		full = map[string]any{"time": start, "value": value}
	}
	if err := json.NewEncoder(w).Encode(full); err != nil {
		panic(err)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status := map[string]interface{}{"status": "ok"}
	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}

}

func LongPooling(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	time.Sleep(6 * time.Second)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("long pooling"))
}
