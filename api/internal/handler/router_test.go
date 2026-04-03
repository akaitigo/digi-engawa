package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akaitigo/digi-engawa/api/internal/handler"
)

func TestHealthEndpoint(t *testing.T) {
	mux := handler.NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	expected := `{"status":"ok"}`
	if rec.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rec.Body.String())
	}
}
