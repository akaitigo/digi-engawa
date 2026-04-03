package service

import (
	"fmt"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/id"
	"github.com/akaitigo/digi-engawa/api/internal/model"
	"github.com/akaitigo/digi-engawa/api/internal/repository"
	"github.com/akaitigo/digi-engawa/api/internal/ws"
)

type ProgressService struct {
	repo          *repository.ProgressRepository
	classroomRepo *repository.ClassroomRepository
	hub           *ws.Hub
}

func NewProgressService(repo *repository.ProgressRepository, classroomRepo *repository.ClassroomRepository, hub *ws.Hub) *ProgressService {
	return &ProgressService{repo: repo, classroomRepo: classroomRepo, hub: hub}
}

func (s *ProgressService) Update(participantID, materialID, classroomID string, currentStep int, completed bool) (model.LearnerProgress, error) {
	existing, found := s.repo.Get(participantID, materialID)
	progressID := ""
	if found {
		progressID = existing.ID
	} else {
		newID, err := id.New()
		if err != nil {
			return model.LearnerProgress{}, fmt.Errorf("generate id: %w", err)
		}
		progressID = newID
	}

	p := model.LearnerProgress{
		ID:            progressID,
		ParticipantID: participantID,
		MaterialID:    materialID,
		CurrentStep:   currentStep,
		Completed:     completed,
		UpdatedAt:     time.Now(),
	}

	if err := s.repo.Upsert(p); err != nil {
		return model.LearnerProgress{}, fmt.Errorf("save progress: %w", err)
	}

	if classroomID != "" {
		s.hub.Broadcast(classroomID, ws.Message{
			Type: "progress_updated",
			Data: p,
		})
	}

	return p, nil
}

func (s *ProgressService) ListByClassroom(classroomID string) []model.LearnerProgress {
	participants := s.classroomRepo.GetParticipants(classroomID)
	ids := make([]string, len(participants))
	for i, p := range participants {
		ids[i] = p.ID
	}
	return s.repo.GetByClassroom(ids)
}
