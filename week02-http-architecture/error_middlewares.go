package main

import (
	"encoding/json"
	"net/http"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				requestID := GetRequestID(r.Context())
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)
				errJson := ErrorJSON{
					RequestID: requestID,
					Error:     "Internal Server Error",
				}
				_ = json.NewEncoder(w).Encode(errJson)
			}
		}()
		next.ServeHTTP(w, r)
	})

}
