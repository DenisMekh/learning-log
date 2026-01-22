package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerError_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/error", nil)
	req = req.WithContext(context.WithValue(req.Context(), RequestIDKey, "test-req-id"))
	rr := httptest.NewRecorder()

	HandlerError(rr, req, ErrNotFound)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status: got %d want %d", rr.Code, http.StatusNotFound)
	}

	if ct := rr.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Fatalf("content-type: got %q", ct)
	}

	var body ErrorJSON
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body.RequestID != "test-req-id" {
		t.Errorf("request id: got %q want %q", body.RequestID, "test-req-id")
	}
	if body.Error == "" {
		t.Errorf("error message empty")
	}
}

func TestHandlerError_AppError(t *testing.T) {
	req := httptest.NewRequest("GET", "/error", nil)
	req = req.WithContext(context.WithValue(req.Context(), RequestIDKey, "rid-123"))
	rr := httptest.NewRecorder()

	err := NewBadRequest(nil, "bad input")
	HandlerError(rr, req, err)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status: got %d want %d", rr.Code, http.StatusBadRequest)
	}

	var body ErrorJSON
	_ = json.NewDecoder(rr.Body).Decode(&body)

	if body.Error != "bad input" {
		t.Errorf("error message mismatch")
	}
}
