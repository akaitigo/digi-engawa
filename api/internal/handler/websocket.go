package handler

import (
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/akaitigo/digi-engawa/api/internal/repository"
	appws "github.com/akaitigo/digi-engawa/api/internal/ws"
)

func newUpgrader() websocket.Upgrader {
	allowedOrigin := os.Getenv("CORS_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:3000"
	}

	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == "" {
				return false
			}
			for _, o := range strings.Split(allowedOrigin, ",") {
				if strings.TrimSpace(o) == origin {
					return true
				}
			}
			return false
		},
	}
}

type WebSocketHandler struct {
	hub           *appws.Hub
	classroomRepo *repository.ClassroomRepository
	upgrader      websocket.Upgrader
}

func NewWebSocketHandler(hub *appws.Hub, classroomRepo *repository.ClassroomRepository) *WebSocketHandler {
	return &WebSocketHandler{
		hub:           hub,
		classroomRepo: classroomRepo,
		upgrader:      newUpgrader(),
	}
}

func (h *WebSocketHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /ws/classroom/{id}", h.handleClassroom)
}

func (h *WebSocketHandler) handleClassroom(w http.ResponseWriter, r *http.Request) {
	classroomID := r.PathValue("id")
	if classroomID == "" {
		http.Error(w, "classroom id required", http.StatusBadRequest)
		return
	}

	if _, ok := h.classroomRepo.GetClassroomByID(classroomID); !ok {
		http.Error(w, "classroom not found", http.StatusNotFound)
		return
	}

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade: %v", err)
		return
	}
	conn.SetReadLimit(4096)

	client := appws.NewClient()
	if err := h.hub.Join(classroomID, client); err != nil {
		log.Printf("websocket join: %v", err)
		_ = conn.Close()
		return
	}

	var closeOnce sync.Once
	cleanup := func() {
		closeOnce.Do(func() {
			h.hub.Leave(classroomID, client)
			client.Close()
			_ = conn.Close()
		})
	}

	go func() {
		defer cleanup()
		for {
			if _, _, readErr := conn.ReadMessage(); readErr != nil {
				return
			}
		}
	}()

	go func() {
		defer cleanup()
		for msg := range client.Messages() {
			if writeErr := conn.WriteMessage(websocket.TextMessage, msg); writeErr != nil {
				return
			}
		}
	}()
}
