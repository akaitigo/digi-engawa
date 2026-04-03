package repository_test

import (
	"testing"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/model"
	"github.com/akaitigo/digi-engawa/api/internal/repository"
)

func TestMaterialRepositorySaveAndGet(t *testing.T) {
	repo, err := repository.NewMaterialRepository(t.TempDir())
	if err != nil {
		t.Fatalf("failed to create repository: %v", err)
	}

	now := time.Now()
	m := model.Material{
		ID:          "test-id-1",
		Title:       "テスト教材",
		Description: "テスト用の教材",
		Steps:       []model.Step{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := repo.Save(m); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	got, ok := repo.GetByID("test-id-1")
	if !ok {
		t.Fatal("expected to find material")
	}

	if got.Title != "テスト教材" {
		t.Errorf("expected title %q, got %q", "テスト教材", got.Title)
	}
}

func TestMaterialRepositoryGetAll(t *testing.T) {
	repo, err := repository.NewMaterialRepository(t.TempDir())
	if err != nil {
		t.Fatalf("failed to create repository: %v", err)
	}

	now := time.Now()
	for i, title := range []string{"教材A", "教材B", "教材C"} {
		m := model.Material{
			ID:        title,
			Title:     title,
			Steps:     []model.Step{},
			CreatedAt: now.Add(time.Duration(i) * time.Second),
			UpdatedAt: now,
		}
		if err := repo.Save(m); err != nil {
			t.Fatalf("failed to save: %v", err)
		}
	}

	all := repo.GetAll()
	if len(all) != 3 {
		t.Errorf("expected 3 materials, got %d", len(all))
	}
}

func TestMaterialRepositoryNotFound(t *testing.T) {
	repo, err := repository.NewMaterialRepository(t.TempDir())
	if err != nil {
		t.Fatalf("failed to create repository: %v", err)
	}

	_, ok := repo.GetByID("nonexistent")
	if ok {
		t.Error("expected not found")
	}
}

func TestMaterialRepositoryPersistence(t *testing.T) {
	dir := t.TempDir()

	repo1, err := repository.NewMaterialRepository(dir)
	if err != nil {
		t.Fatalf("failed to create repository: %v", err)
	}

	now := time.Now()
	m := model.Material{
		ID:        "persist-test",
		Title:     "永続化テスト",
		Steps:     []model.Step{},
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := repo1.Save(m); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	// Create new repo instance from same dir
	repo2, err := repository.NewMaterialRepository(dir)
	if err != nil {
		t.Fatalf("failed to create second repository: %v", err)
	}

	got, ok := repo2.GetByID("persist-test")
	if !ok {
		t.Fatal("expected to find material after reload")
	}
	if got.Title != "永続化テスト" {
		t.Errorf("expected title %q, got %q", "永続化テスト", got.Title)
	}
}
