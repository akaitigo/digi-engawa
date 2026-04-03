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

func setupRouter(t *testing.T) *http.ServeMux {
	t.Helper()
	mux, err := handler.NewRouter(t.TempDir())
	if err != nil {
		t.Fatalf("failed to create router: %v", err)
	}
	return mux
}

func TestCreateMaterial(t *testing.T) {
	mux := setupRouter(t)

	body := `{"title":"スマホの使い方","description":"初心者向けスマホ教材"}`
	req := httptest.NewRequest(http.MethodPost, "/api/materials", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, rec.Code, rec.Body.String())
	}

	var m model.Material
	if err := json.NewDecoder(rec.Body).Decode(&m); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if m.Title != "スマホの使い方" {
		t.Errorf("expected title %q, got %q", "スマホの使い方", m.Title)
	}
	if m.ID == "" {
		t.Error("expected non-empty ID")
	}
}

func TestCreateMaterialValidation(t *testing.T) {
	mux := setupRouter(t)

	body := `{"title":"","description":"no title"}`
	req := httptest.NewRequest(http.MethodPost, "/api/materials", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestListMaterials(t *testing.T) {
	mux := setupRouter(t)

	// Create a material first
	body := `{"title":"テスト教材","description":"テスト"}`
	req := httptest.NewRequest(http.MethodPost, "/api/materials", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	// List materials
	req = httptest.NewRequest(http.MethodGet, "/api/materials", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var materials []model.Material
	if err := json.NewDecoder(rec.Body).Decode(&materials); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(materials) != 1 {
		t.Errorf("expected 1 material, got %d", len(materials))
	}
}

func TestGetMaterial(t *testing.T) {
	mux := setupRouter(t)

	// Create
	body := `{"title":"テスト教材","description":"テスト"}`
	req := httptest.NewRequest(http.MethodPost, "/api/materials", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var created model.Material
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	// Get
	req = httptest.NewRequest(http.MethodGet, "/api/materials/"+created.ID, nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var m model.Material
	if err := json.NewDecoder(rec.Body).Decode(&m); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	if m.Title != "テスト教材" {
		t.Errorf("expected title %q, got %q", "テスト教材", m.Title)
	}
}

func TestGetMaterialNotFound(t *testing.T) {
	mux := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/api/materials/nonexistent", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestAddStepAndGetStep(t *testing.T) {
	mux := setupRouter(t)

	// Create material
	body := `{"title":"テスト教材","description":"テスト"}`
	req := httptest.NewRequest(http.MethodPost, "/api/materials", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var created model.Material
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	// Add step
	stepBody := `{"title":"電源を入れる","body":"スマホの横にあるボタンを長く押します","furigana_body":"スマホのよこにあるボタンをながくおします","audio_text":"スマホの横にあるボタンを長く押してください"}`
	req = httptest.NewRequest(http.MethodPost, "/api/materials/"+created.ID+"/steps", bytes.NewBufferString(stepBody))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, rec.Code, rec.Body.String())
	}

	var step model.Step
	if err := json.NewDecoder(rec.Body).Decode(&step); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	if step.StepOrder != 1 {
		t.Errorf("expected step_order 1, got %d", step.StepOrder)
	}

	// Get step
	req = httptest.NewRequest(http.MethodGet, "/api/materials/"+created.ID+"/steps/1", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var gotStep model.Step
	if err := json.NewDecoder(rec.Body).Decode(&gotStep); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	if gotStep.Title != "電源を入れる" {
		t.Errorf("expected title %q, got %q", "電源を入れる", gotStep.Title)
	}
	if gotStep.FuriganaBody == "" {
		t.Error("expected non-empty furigana_body")
	}
}

func TestAddStepValidation(t *testing.T) {
	mux := setupRouter(t)

	// Create material
	body := `{"title":"テスト教材","description":"テスト"}`
	req := httptest.NewRequest(http.MethodPost, "/api/materials", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var created model.Material
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	// Add step without body
	stepBody := `{"title":"電源を入れる","body":""}`
	req = httptest.NewRequest(http.MethodPost, "/api/materials/"+created.ID+"/steps", bytes.NewBufferString(stepBody))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
