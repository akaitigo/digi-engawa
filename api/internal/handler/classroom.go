package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/service"
)

type ClassroomHandler struct {
	svc *service.ClassroomService
}

func NewClassroomHandler(svc *service.ClassroomService) *ClassroomHandler {
	return &ClassroomHandler{svc: svc}
}

func (h *ClassroomHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/classrooms", h.handleCreate)
	mux.HandleFunc("GET /api/classrooms", h.handleList)
	mux.HandleFunc("GET /api/classrooms/{id}", h.handleGet)
	mux.HandleFunc("POST /api/classrooms/{id}/participants", h.handleAddParticipant)
	mux.HandleFunc("GET /api/classrooms/{id}/participants", h.handleListParticipants)
	mux.HandleFunc("GET /api/join/{code}", h.handleJoinByCode)
}

func (h *ClassroomHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Location    string `json:"location"`
		Capacity    int    `json:"capacity"`
		ScheduledAt string `json:"scheduled_at"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.Title == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "title is required"})
		return
	}

	scheduledAt, err := time.Parse(time.RFC3339, req.ScheduledAt)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid scheduled_at format (use RFC3339)"})
		return
	}

	if req.Capacity <= 0 {
		req.Capacity = 20
	}

	c, err := h.svc.Create(req.Title, req.Description, req.Location, req.Capacity, scheduledAt)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

func (h *ClassroomHandler) handleList(w http.ResponseWriter, _ *http.Request) {
	classrooms := h.svc.List()
	writeJSON(w, http.StatusOK, classrooms)
}

func (h *ClassroomHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	c, err := h.svc.Get(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (h *ClassroomHandler) handleAddParticipant(w http.ResponseWriter, r *http.Request) {
	classroomID := r.PathValue("id")

	var req struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.Name == "" || req.Role == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name and role are required"})
		return
	}

	p, err := h.svc.AddParticipant(classroomID, req.Name, req.Role)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}
	writeJSON(w, http.StatusCreated, p)
}

func (h *ClassroomHandler) handleListParticipants(w http.ResponseWriter, r *http.Request) {
	classroomID := r.PathValue("id")
	participants := h.svc.ListParticipants(classroomID)
	writeJSON(w, http.StatusOK, participants)
}

func (h *ClassroomHandler) handleJoinByCode(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	c, err := h.svc.GetByCode(code)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, c)
}
