package handler

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/model"
	"github.com/akaitigo/digi-engawa/api/internal/repository"
	"github.com/akaitigo/digi-engawa/api/internal/ws"
)

type ProgressHandler struct {
	repo          *repository.ProgressRepository
	classroomRepo *repository.ClassroomRepository
	hub           *ws.Hub
}

func NewProgressHandler(repo *repository.ProgressRepository, classroomRepo *repository.ClassroomRepository, hub *ws.Hub) *ProgressHandler {
	return &ProgressHandler{repo: repo, classroomRepo: classroomRepo, hub: hub}
}

func (h *ProgressHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("PUT /api/progress", h.handleUpdate)
	mux.HandleFunc("GET /api/classrooms/{id}/progress", h.handleGetByClassroom)
}

func (h *ProgressHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ParticipantID string `json:"participant_id"`
		MaterialID    string `json:"material_id"`
		CurrentStep   int    `json:"current_step"`
		Completed     bool   `json:"completed"`
		ClassroomID   string `json:"classroom_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.ParticipantID == "" || req.MaterialID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "participant_id and material_id are required"})
		return
	}

	existing, found := h.repo.Get(req.ParticipantID, req.MaterialID)
	id := newID()
	if found {
		id = existing.ID
	}

	p := model.LearnerProgress{
		ID:            id,
		ParticipantID: req.ParticipantID,
		MaterialID:    req.MaterialID,
		CurrentStep:   req.CurrentStep,
		Completed:     req.Completed,
		UpdatedAt:     time.Now(),
	}

	if err := h.repo.Upsert(p); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if req.ClassroomID != "" {
		h.hub.Broadcast(req.ClassroomID, ws.Message{
			Type: "progress_updated",
			Data: p,
		})
	}

	writeJSON(w, http.StatusOK, p)
}

func (h *ProgressHandler) handleGetByClassroom(w http.ResponseWriter, r *http.Request) {
	classroomID := r.PathValue("id")

	participants := h.classroomRepo.GetParticipants(classroomID)
	ids := make([]string, len(participants))
	for i, p := range participants {
		ids[i] = p.ID
	}

	progress := h.repo.GetByClassroom(ids)
	writeJSON(w, http.StatusOK, progress)
}

func newID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("failed to generate ID: %v", err))
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
