package handler

import (
	"encoding/json"
	"net/http"

	"github.com/akaitigo/digi-engawa/api/internal/service"
)

type HelpRequestHandler struct {
	svc *service.HelpRequestService
}

func NewHelpRequestHandler(svc *service.HelpRequestService) *HelpRequestHandler {
	return &HelpRequestHandler{svc: svc}
}

func (h *HelpRequestHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/help-requests", h.handleCreate)
	mux.HandleFunc("PATCH /api/help-requests/{id}", h.handleUpdateStatus)
	mux.HandleFunc("GET /api/classrooms/{id}/help-requests", h.handleListByClassroom)
}

func (h *HelpRequestHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ClassroomID    string `json:"classroom_id"`
		ParticipantID  string `json:"participant_id"`
		MaterialStepID string `json:"material_step_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.ClassroomID == "" || req.ParticipantID == "" || req.MaterialStepID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "classroom_id, participant_id, and material_step_id are required"})
		return
	}

	hr, err := h.svc.Create(req.ClassroomID, req.ParticipantID, req.MaterialStepID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}
	writeJSON(w, http.StatusCreated, hr)
}

func (h *HelpRequestHandler) handleUpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.Status == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "status is required"})
		return
	}

	hr, err := h.svc.UpdateStatus(id, req.Status)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}
	writeJSON(w, http.StatusOK, hr)
}

func (h *HelpRequestHandler) handleListByClassroom(w http.ResponseWriter, r *http.Request) {
	classroomID := r.PathValue("id")
	requests := h.svc.ListByClassroom(classroomID)
	writeJSON(w, http.StatusOK, requests)
}
