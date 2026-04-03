package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akaitigo/digi-engawa/api/internal/handler"
	"github.com/akaitigo/digi-engawa/api/internal/model"
)

func TestCreateHelpRequest(t *testing.T) {
	mux := setupRouter(t)

	body := `{"classroom_id":"class-1","participant_id":"part-1","material_step_id":"step-1"}`
	req := httptest.NewRequest(http.MethodPost, "/api/help-requests", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d: %s", http.StatusCreated, rec.Code, rec.Body.String())
	}

	var hr model.HelpRequest
	if err := json.NewDecoder(rec.Body).Decode(&hr); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if hr.Status != "pending" {
		t.Errorf("expected status pending, got %q", hr.Status)
	}
	if hr.ClassroomID != "class-1" {
		t.Errorf("expected classroom_id class-1, got %q", hr.ClassroomID)
	}
}

func TestCreateHelpRequestValidation(t *testing.T) {
	mux := setupRouter(t)

	body := `{"classroom_id":"","participant_id":"part-1","material_step_id":"step-1"}`
	req := httptest.NewRequest(http.MethodPost, "/api/help-requests", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateHelpRequestStatus(t *testing.T) {
	mux := setupRouter(t)

	// Create
	body := `{"classroom_id":"class-1","participant_id":"part-1","material_step_id":"step-1"}`
	req := httptest.NewRequest(http.MethodPost, "/api/help-requests", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var created model.HelpRequest
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatalf("decode: %v", err)
	}

	// Update to in_progress
	updateBody := `{"status":"in_progress"}`
	req = httptest.NewRequest(http.MethodPatch, "/api/help-requests/"+created.ID, bytes.NewBufferString(updateBody))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}

	var updated model.HelpRequest
	if err := json.NewDecoder(rec.Body).Decode(&updated); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if updated.Status != "in_progress" {
		t.Errorf("expected in_progress, got %q", updated.Status)
	}
}

func TestUpdateHelpRequestInvalidTransition(t *testing.T) {
	mux := setupRouter(t)

	// Create
	body := `{"classroom_id":"class-1","participant_id":"part-1","material_step_id":"step-1"}`
	req := httptest.NewRequest(http.MethodPost, "/api/help-requests", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var created model.HelpRequest
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatalf("decode: %v", err)
	}

	// Try to go directly to resolved (invalid: should go through in_progress first)
	updateBody := `{"status":"resolved"}`
	req = httptest.NewRequest(http.MethodPatch, "/api/help-requests/"+created.ID, bytes.NewBufferString(updateBody))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected %d for invalid transition, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestListHelpRequestsByClassroom(t *testing.T) {
	mux := setupRouter(t)

	// Create 2 requests for class-1
	for i := 0; i < 2; i++ {
		body := `{"classroom_id":"class-1","participant_id":"part-1","material_step_id":"step-1"}`
		req := httptest.NewRequest(http.MethodPost, "/api/help-requests", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
	}

	// Create 1 for class-2
	body := `{"classroom_id":"class-2","participant_id":"part-2","material_step_id":"step-1"}`
	req := httptest.NewRequest(http.MethodPost, "/api/help-requests", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	// List for class-1
	req = httptest.NewRequest(http.MethodGet, "/api/classrooms/class-1/help-requests", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}

	var requests []model.HelpRequest
	if err := json.NewDecoder(rec.Body).Decode(&requests); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if len(requests) != 2 {
		t.Errorf("expected 2, got %d", len(requests))
	}
}

// setupRouter is defined in material_test.go via helper
func init() {
	_ = handler.NewRouter
}
