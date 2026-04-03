package handler

import (
	"encoding/json"
	"net/http"

	"github.com/akaitigo/digi-engawa/api/internal/service"
)

type ProgressHandler struct {
	svc *service.ProgressService
}

func NewProgressHandler(svc *service.ProgressService) *ProgressHandler {
	return &ProgressHandler{svc: svc}
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

	p, err := h.svc.Update(req.ParticipantID, req.MaterialID, req.ClassroomID, req.CurrentStep, req.Completed)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	writeJSON(w, http.StatusOK, p)
}

func (h *ProgressHandler) handleGetByClassroom(w http.ResponseWriter, r *http.Request) {
	classroomID := r.PathValue("id")
	progress := h.svc.ListByClassroom(classroomID)
	writeJSON(w, http.StatusOK, progress)
}
