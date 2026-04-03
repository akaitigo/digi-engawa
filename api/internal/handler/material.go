package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/akaitigo/digi-engawa/api/internal/service"
)

type MaterialHandler struct {
	svc *service.MaterialService
}

func NewMaterialHandler(svc *service.MaterialService) *MaterialHandler {
	return &MaterialHandler{svc: svc}
}

func (h *MaterialHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/materials", h.handleList)
	mux.HandleFunc("GET /api/materials/{id}", h.handleGet)
	mux.HandleFunc("GET /api/materials/{id}/steps/{order}", h.handleGetStep)
	mux.HandleFunc("POST /api/materials", h.handleCreate)
	mux.HandleFunc("POST /api/materials/{id}/steps", h.handleAddStep)
}

func (h *MaterialHandler) handleList(w http.ResponseWriter, _ *http.Request) {
	materials := h.svc.ListMaterials()
	writeJSON(w, http.StatusOK, materials)
}

func (h *MaterialHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	m, err := h.svc.GetMaterial(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, m)
}

func (h *MaterialHandler) handleGetStep(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	orderStr := r.PathValue("order")

	order, err := strconv.Atoi(orderStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid step order"})
		return
	}

	step, err := h.svc.GetStep(id, order)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, step)
}

func (h *MaterialHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.Title == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "title is required"})
		return
	}

	m, err := h.svc.CreateMaterial(req.Title, req.Description)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}
	writeJSON(w, http.StatusCreated, m)
}

func (h *MaterialHandler) handleAddStep(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req struct {
		Title        string `json:"title"`
		Body         string `json:"body"`
		FuriganaBody string `json:"furigana_body"`
		AudioText    string `json:"audio_text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.Title == "" || req.Body == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "title and body are required"})
		return
	}

	step, err := h.svc.AddStep(id, req.Title, req.Body, req.FuriganaBody, req.AudioText)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}
	writeJSON(w, http.StatusCreated, step)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
