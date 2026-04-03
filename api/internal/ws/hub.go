package ws

import (
	"encoding/json"
	"errors"
	"sync"
)

const (
	MaxClientsPerRoom = 100
	MaxTotalClients   = 1000
)

var ErrRoomFull = errors.New("room connection limit reached")

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Hub struct {
	mu          sync.RWMutex
	rooms       map[string]map[*Client]bool
	totalCount  int
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Join(roomID string, client *Client) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.totalCount >= MaxTotalClients {
		return ErrRoomFull
	}

	if h.rooms[roomID] == nil {
		h.rooms[roomID] = make(map[*Client]bool)
	}
	if len(h.rooms[roomID]) >= MaxClientsPerRoom {
		return ErrRoomFull
	}

	h.rooms[roomID][client] = true
	h.totalCount++
	return nil
}

func (h *Hub) Leave(roomID string, client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if clients, ok := h.rooms[roomID]; ok {
		if clients[client] {
			delete(clients, client)
			h.totalCount--
		}
		if len(clients) == 0 {
			delete(h.rooms, roomID)
		}
	}
}

func (h *Hub) Broadcast(roomID string, msg Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		return
	}

	h.mu.RLock()
	clients := make([]*Client, 0, len(h.rooms[roomID]))
	for client := range h.rooms[roomID] {
		clients = append(clients, client)
	}
	h.mu.RUnlock()

	for _, client := range clients {
		client.Send(data)
	}
}

func (h *Hub) RoomCount(roomID string) int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if clients, ok := h.rooms[roomID]; ok {
		return len(clients)
	}
	return 0
}
