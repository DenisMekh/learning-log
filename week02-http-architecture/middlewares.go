package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// RequestIDMiddleware мидлварь для добавления X-Request-Id в контекст
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), "requestID", requestID)
		w.Header().Add("X-Request-Id", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// LoggingMiddleware для логгирования HTTP метод/URL/статус код/время выполнения/request-id
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID := r.Context().Value("requestID").(string)
		url := r.URL.String()
		next.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s\t%s", r.Method, url, r.Response.StatusCode, time.Since(start), requestID)
	})
}
