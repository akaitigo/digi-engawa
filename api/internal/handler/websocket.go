package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	appws "github.com/akaitigo/digi-engawa/api/internal/ws"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: restrict to allowed origins in production
		return true
	},
}

type WebSocketHandler struct {
	hub *appws.Hub
}

func NewWebSocketHandler(hub *appws.Hub) *WebSocketHandler {
	return &WebSocketHandler{hub: hub}
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

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade: %v", err)
		return
	}

	client := appws.NewClient()
	h.hub.Join(classroomID, client)

	go func() {
		defer func() {
			h.hub.Leave(classroomID, client)
			client.Close()
			conn.Close()
		}()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()

	go func() {
		defer conn.Close()
		for msg := range client.Messages() {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		}
	}()
}
