package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akaitigo/digi-engawa/api/internal/model"
)

func createClassroom(t *testing.T, mux *http.ServeMux) model.Classroom {
	t.Helper()
	body := `{"title":"テスト教室","description":"テスト","location":"公民館A","capacity":10,"scheduled_at":"2026-05-01T10:00:00Z"}`
	req := httptest.NewRequest(http.MethodPost, "/api/classrooms", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("create classroom: expected %d, got %d: %s", http.StatusCreated, rec.Code, rec.Body.String())
	}

	var c model.Classroom
	if err := json.NewDecoder(rec.Body).Decode(&c); err != nil {
		t.Fatalf("decode: %v", err)
	}
	return c
}

func TestCreateClassroom(t *testing.T) {
	mux := setupRouter(t)
	c := createClassroom(t, mux)

	if c.Title != "テスト教室" {
		t.Errorf("expected title テスト教室, got %q", c.Title)
	}
	if c.ClassroomCode == "" {
		t.Error("expected non-empty classroom code")
	}
	if len(c.ClassroomCode) != 6 {
		t.Errorf("expected 6-char code, got %d", len(c.ClassroomCode))
	}
}

func TestCreateClassroomValidation(t *testing.T) {
	mux := setupRouter(t)

	body := `{"title":"","scheduled_at":"2026-05-01T10:00:00Z"}`
	req := httptest.NewRequest(http.MethodPost, "/api/classrooms", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestListClassrooms(t *testing.T) {
	mux := setupRouter(t)
	createClassroom(t, mux)
	createClassroom(t, mux)

	req := httptest.NewRequest(http.MethodGet, "/api/classrooms", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}

	var classrooms []model.Classroom
	if err := json.NewDecoder(rec.Body).Decode(&classrooms); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if len(classrooms) != 2 {
		t.Errorf("expected 2, got %d", len(classrooms))
	}
}

func TestGetClassroom(t *testing.T) {
	mux := setupRouter(t)
	created := createClassroom(t, mux)

	req := httptest.NewRequest(http.MethodGet, "/api/classrooms/"+created.ID, nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}

	var c model.Classroom
	if err := json.NewDecoder(rec.Body).Decode(&c); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if c.Title != "テスト教室" {
		t.Errorf("expected テスト教室, got %q", c.Title)
	}
}

func TestJoinByCode(t *testing.T) {
	mux := setupRouter(t)
	created := createClassroom(t, mux)

	req := httptest.NewRequest(http.MethodGet, "/api/join/"+created.ClassroomCode, nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}

	var c model.Classroom
	if err := json.NewDecoder(rec.Body).Decode(&c); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if c.ID != created.ID {
		t.Errorf("expected ID %q, got %q", created.ID, c.ID)
	}
}

func TestAddAndListParticipants(t *testing.T) {
	mux := setupRouter(t)
	created := createClassroom(t, mux)

	// Add participants
	for _, p := range []struct{ name, role string }{
		{"田中太郎", "learner"},
		{"山田花子", "supporter"},
		{"佐藤一郎", "organizer"},
	} {
		body, _ := json.Marshal(map[string]string{"name": p.name, "role": p.role})
		req := httptest.NewRequest(http.MethodPost, "/api/classrooms/"+created.ID+"/participants", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Fatalf("add participant: expected %d, got %d: %s", http.StatusCreated, rec.Code, rec.Body.String())
		}
	}

	// List
	req := httptest.NewRequest(http.MethodGet, "/api/classrooms/"+created.ID+"/participants", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}

	var participants []model.Participant
	if err := json.NewDecoder(rec.Body).Decode(&participants); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if len(participants) != 3 {
		t.Errorf("expected 3, got %d", len(participants))
	}
}

func TestAddParticipantInvalidRole(t *testing.T) {
	mux := setupRouter(t)
	created := createClassroom(t, mux)

	body := `{"name":"テスト","role":"admin"}`
	req := httptest.NewRequest(http.MethodPost, "/api/classrooms/"+created.ID+"/participants", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected %d for invalid role, got %d", http.StatusBadRequest, rec.Code)
	}
}
