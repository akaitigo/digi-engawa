package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/id"
	"github.com/akaitigo/digi-engawa/api/internal/model"
	"github.com/akaitigo/digi-engawa/api/internal/repository"
)

type ClassroomService struct {
	repo *repository.ClassroomRepository
}

func NewClassroomService(repo *repository.ClassroomRepository) *ClassroomService {
	return &ClassroomService{repo: repo}
}

func (s *ClassroomService) Create(title, description, location string, capacity int, scheduledAt time.Time) (model.Classroom, error) {
	newID, err := id.New()
	if err != nil {
		return model.Classroom{}, fmt.Errorf("generate id: %w", err)
	}

	code, err := generateClassroomCode()
	if err != nil {
		return model.Classroom{}, fmt.Errorf("generate code: %w", err)
	}

	now := time.Now()
	c := model.Classroom{
		ID:            newID,
		Title:         title,
		Description:   description,
		Location:      location,
		Capacity:      capacity,
		ScheduledAt:   scheduledAt,
		ClassroomCode: code,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := s.repo.SaveClassroom(c); err != nil {
		return model.Classroom{}, fmt.Errorf("save classroom: %w", err)
	}

	return c, nil
}

func (s *ClassroomService) List() []model.Classroom {
	return s.repo.GetAllClassrooms()
}

func (s *ClassroomService) Get(classroomID string) (model.Classroom, error) {
	c, ok := s.repo.GetClassroomByID(classroomID)
	if !ok {
		return model.Classroom{}, fmt.Errorf("classroom not found: %s", classroomID)
	}
	return c, nil
}

func (s *ClassroomService) GetByCode(code string) (model.Classroom, error) {
	c, ok := s.repo.GetClassroomByCode(code)
	if !ok {
		return model.Classroom{}, fmt.Errorf("classroom not found for code: %s", code)
	}
	return c, nil
}

func (s *ClassroomService) AddParticipant(classroomID, name, role string) (model.Participant, error) {
	if _, ok := s.repo.GetClassroomByID(classroomID); !ok {
		return model.Participant{}, fmt.Errorf("classroom not found: %s", classroomID)
	}

	validRoles := map[string]bool{"learner": true, "supporter": true, "organizer": true}
	if !validRoles[role] {
		return model.Participant{}, fmt.Errorf("invalid role: %s (must be learner, supporter, or organizer)", role)
	}

	partID, err := id.New()
	if err != nil {
		return model.Participant{}, fmt.Errorf("generate id: %w", err)
	}

	p := model.Participant{
		ID:          partID,
		ClassroomID: classroomID,
		Name:        name,
		Role:        role,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.AddParticipant(p); err != nil {
		return model.Participant{}, fmt.Errorf("save participant: %w", err)
	}

	return p, nil
}

func (s *ClassroomService) ListParticipants(classroomID string) []model.Participant {
	return s.repo.GetParticipants(classroomID)
}

func generateClassroomCode() (string, error) {
	const charset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	code := make([]byte, 6)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[n.Int64()]
	}
	return string(code), nil
}
