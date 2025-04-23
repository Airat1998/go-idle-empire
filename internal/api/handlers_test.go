package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClick(t *testing.T) {
	req := httptest.NewRequest("POST", "/click", nil)
	w := httptest.NewRecorder()
	HandleClick(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestStatus(t *testing.T) {
	req := httptest.NewRequest("GET", "/status", nil)
	w := httptest.NewRecorder()
	HandleStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
