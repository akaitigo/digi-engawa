package service_test

import (
	"testing"

	"github.com/akaitigo/digi-engawa/api/internal/repository"
	"github.com/akaitigo/digi-engawa/api/internal/service"
)

func setupService(t *testing.T) *service.MaterialService {
	t.Helper()
	repo, err := repository.NewMaterialRepository(t.TempDir())
	if err != nil {
		t.Fatalf("failed to create repository: %v", err)
	}
	return service.NewMaterialService(repo)
}

func TestCreateAndGetMaterial(t *testing.T) {
	svc := setupService(t)

	m, err := svc.CreateMaterial("テスト教材", "説明")
	if err != nil {
		t.Fatalf("failed to create: %v", err)
	}

	got, err := svc.GetMaterial(m.ID)
	if err != nil {
		t.Fatalf("failed to get: %v", err)
	}

	if got.Title != "テスト教材" {
		t.Errorf("expected title %q, got %q", "テスト教材", got.Title)
	}
}

func TestGetMaterialNotFound(t *testing.T) {
	svc := setupService(t)

	_, err := svc.GetMaterial("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent material")
	}
}

func TestAddStepAndGetStep(t *testing.T) {
	svc := setupService(t)

	m, err := svc.CreateMaterial("テスト教材", "説明")
	if err != nil {
		t.Fatalf("failed to create: %v", err)
	}

	step, err := svc.AddStep(m.ID, "ステップ1", "内容", "ないよう", "音声テキスト")
	if err != nil {
		t.Fatalf("failed to add step: %v", err)
	}

	if step.StepOrder != 1 {
		t.Errorf("expected step_order 1, got %d", step.StepOrder)
	}

	got, err := svc.GetStep(m.ID, 1)
	if err != nil {
		t.Fatalf("failed to get step: %v", err)
	}

	if got.Title != "ステップ1" {
		t.Errorf("expected title %q, got %q", "ステップ1", got.Title)
	}
}

func TestGetStepNotFound(t *testing.T) {
	svc := setupService(t)

	m, err := svc.CreateMaterial("テスト教材", "説明")
	if err != nil {
		t.Fatalf("failed to create: %v", err)
	}

	_, err = svc.GetStep(m.ID, 999)
	if err == nil {
		t.Error("expected error for nonexistent step")
	}
}

func TestListMaterials(t *testing.T) {
	svc := setupService(t)

	_, err := svc.CreateMaterial("教材1", "")
	if err != nil {
		t.Fatalf("failed to create: %v", err)
	}
	_, err = svc.CreateMaterial("教材2", "")
	if err != nil {
		t.Fatalf("failed to create: %v", err)
	}

	list := svc.ListMaterials()
	if len(list) != 2 {
		t.Errorf("expected 2 materials, got %d", len(list))
	}
}
