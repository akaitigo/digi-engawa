package service

import (
	"fmt"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/id"
	"github.com/akaitigo/digi-engawa/api/internal/model"
	"github.com/akaitigo/digi-engawa/api/internal/repository"
	"github.com/akaitigo/digi-engawa/api/internal/ws"
)

type HelpRequestService struct {
	repo *repository.HelpRequestRepository
	hub  *ws.Hub
}

func NewHelpRequestService(repo *repository.HelpRequestRepository, hub *ws.Hub) *HelpRequestService {
	return &HelpRequestService{repo: repo, hub: hub}
}

func (s *HelpRequestService) Create(classroomID, participantID, materialStepID string) (model.HelpRequest, error) {
	newID, err := id.New()
	if err != nil {
		return model.HelpRequest{}, fmt.Errorf("generate id: %w", err)
	}

	now := time.Now()
	hr := model.HelpRequest{
		ID:             newID,
		ClassroomID:    classroomID,
		ParticipantID:  participantID,
		MaterialStepID: materialStepID,
		Status:         "pending",
		CreatedAt:      now,
	}

	if err := s.repo.Save(hr); err != nil {
		return model.HelpRequest{}, fmt.Errorf("save help request: %w", err)
	}

	s.hub.Broadcast(classroomID, ws.Message{
		Type: "help_request_created",
		Data: hr,
	})

	return hr, nil
}

func (s *HelpRequestService) UpdateStatus(requestID, status string) (model.HelpRequest, error) {
	hr, ok := s.repo.GetByID(requestID)
	if !ok {
		return model.HelpRequest{}, fmt.Errorf("help request not found: %s", requestID)
	}

	validTransitions := map[string][]string{
		"pending":     {"in_progress"},
		"in_progress": {"resolved"},
	}

	allowed, exists := validTransitions[hr.Status]
	if !exists {
		return model.HelpRequest{}, fmt.Errorf("cannot transition from status %q", hr.Status)
	}

	valid := false
	for _, s := range allowed {
		if s == status {
			valid = true
			break
		}
	}
	if !valid {
		return model.HelpRequest{}, fmt.Errorf("invalid transition from %q to %q", hr.Status, status)
	}

	hr.Status = status
	if status == "resolved" {
		now := time.Now()
		hr.ResolvedAt = &now
	}

	if err := s.repo.Save(hr); err != nil {
		return model.HelpRequest{}, fmt.Errorf("save help request: %w", err)
	}

	s.hub.Broadcast(hr.ClassroomID, ws.Message{
		Type: "help_request_updated",
		Data: hr,
	})

	return hr, nil
}

func (s *HelpRequestService) ListByClassroom(classroomID string) []model.HelpRequest {
	return s.repo.GetByClassroom(classroomID)
}
