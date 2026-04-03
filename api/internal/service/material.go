package service

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/model"
	"github.com/akaitigo/digi-engawa/api/internal/repository"
)

type MaterialService struct {
	repo *repository.MaterialRepository
}

func NewMaterialService(repo *repository.MaterialRepository) *MaterialService {
	return &MaterialService{repo: repo}
}

func (s *MaterialService) ListMaterials() []model.Material {
	return s.repo.GetAll()
}

func (s *MaterialService) GetMaterial(id string) (model.Material, error) {
	m, ok := s.repo.GetByID(id)
	if !ok {
		return model.Material{}, fmt.Errorf("material not found: %s", id)
	}
	return m, nil
}

func (s *MaterialService) GetStep(materialID string, stepOrder int) (model.Step, error) {
	m, ok := s.repo.GetByID(materialID)
	if !ok {
		return model.Step{}, fmt.Errorf("material not found: %s", materialID)
	}

	for _, step := range m.Steps {
		if step.StepOrder == stepOrder {
			return step, nil
		}
	}

	return model.Step{}, fmt.Errorf("step %d not found in material %s", stepOrder, materialID)
}

func (s *MaterialService) CreateMaterial(title, description string) (model.Material, error) {
	now := time.Now()
	m := model.Material{
		ID:          generateID(),
		Title:       title,
		Description: description,
		Steps:       []model.Step{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Save(m); err != nil {
		return model.Material{}, fmt.Errorf("save material: %w", err)
	}

	return m, nil
}

func (s *MaterialService) AddStep(materialID string, title, body, furiganaBody, audioText string) (model.Step, error) {
	m, ok := s.repo.GetByID(materialID)
	if !ok {
		return model.Step{}, fmt.Errorf("material not found: %s", materialID)
	}

	step := model.Step{
		ID:           generateID(),
		MaterialID:   materialID,
		StepOrder:    len(m.Steps) + 1,
		Title:        title,
		Body:         body,
		FuriganaBody: furiganaBody,
		AudioText:    audioText,
		CreatedAt:    time.Now(),
	}

	m.Steps = append(m.Steps, step)
	m.UpdatedAt = time.Now()

	if err := s.repo.Save(m); err != nil {
		return model.Step{}, fmt.Errorf("save step: %w", err)
	}

	return step, nil
}

func generateID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("failed to generate ID: %v", err))
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
