package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RequestIDType struct{}

var RequestIDKey = RequestIDType{}

type FakeWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *FakeWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// RequestIDMiddleware мидлварь для добавления X-Request-Id в контекст
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		w.Header().Add("X-Request-Id", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// LoggingMiddleware для логгирования HTTP метод/URL/статус код/время выполнения/request-id
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fw := &FakeWriter{ResponseWriter: w, statusCode: http.StatusOK}
		start := time.Now()
		requestID := GetRequestID(r.Context())
		url := r.URL.String()
		next.ServeHTTP(fw, r)
		log.Printf("%s\t%s\t%d\t%s\t%s", r.Method, url, fw.statusCode, time.Since(start), requestID)
	})
}

func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}
	return ""
}
