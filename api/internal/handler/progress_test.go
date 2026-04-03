package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akaitigo/digi-engawa/api/internal/model"
)

func TestUpdateProgress(t *testing.T) {
	mux := setupRouter(t)

	body := `{"participant_id":"part-1","material_id":"mat-1","current_step":3,"completed":false}`
	req := httptest.NewRequest(http.MethodPut, "/api/progress", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}

	var p model.LearnerProgress
	if err := json.NewDecoder(rec.Body).Decode(&p); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if p.CurrentStep != 3 {
		t.Errorf("expected step 3, got %d", p.CurrentStep)
	}
	if p.Completed {
		t.Error("expected not completed")
	}
}

func TestUpdateProgressUpsert(t *testing.T) {
	mux := setupRouter(t)

	// Create progress
	body := `{"participant_id":"part-1","material_id":"mat-1","current_step":1,"completed":false}`
	req := httptest.NewRequest(http.MethodPut, "/api/progress", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var first model.LearnerProgress
	if err := json.NewDecoder(rec.Body).Decode(&first); err != nil {
		t.Fatalf("decode: %v", err)
	}

	// Update same participant+material
	body = `{"participant_id":"part-1","material_id":"mat-1","current_step":5,"completed":true}`
	req = httptest.NewRequest(http.MethodPut, "/api/progress", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var second model.LearnerProgress
	if err := json.NewDecoder(rec.Body).Decode(&second); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if second.ID != first.ID {
		t.Errorf("expected same ID on upsert, got %q vs %q", first.ID, second.ID)
	}
	if second.CurrentStep != 5 {
		t.Errorf("expected step 5, got %d", second.CurrentStep)
	}
	if !second.Completed {
		t.Error("expected completed")
	}
}

func TestUpdateProgressValidation(t *testing.T) {
	mux := setupRouter(t)

	body := `{"participant_id":"","material_id":"mat-1","current_step":1}`
	req := httptest.NewRequest(http.MethodPut, "/api/progress", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestGetProgressByClassroom(t *testing.T) {
	mux := setupRouter(t)

	// Create classroom and participant
	classroom := createClassroom(t, mux)
	partBody := `{"name":"テスト受講者","role":"learner"}`
	req := httptest.NewRequest(http.MethodPost, "/api/classrooms/"+classroom.ID+"/participants", bytes.NewBufferString(partBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var participant model.Participant
	if err := json.NewDecoder(rec.Body).Decode(&participant); err != nil {
		t.Fatalf("decode: %v", err)
	}

	// Update progress for participant
	progressBody, _ := json.Marshal(map[string]interface{}{
		"participant_id": participant.ID,
		"material_id":    "mat-1",
		"current_step":   3,
		"completed":      false,
	})
	req = httptest.NewRequest(http.MethodPut, "/api/progress", bytes.NewBuffer(progressBody))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	// Get classroom progress
	req = httptest.NewRequest(http.MethodGet, "/api/classrooms/"+classroom.ID+"/progress", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}

	var progress []model.LearnerProgress
	if err := json.NewDecoder(rec.Body).Decode(&progress); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if len(progress) != 1 {
		t.Errorf("expected 1 progress, got %d", len(progress))
	}
}
