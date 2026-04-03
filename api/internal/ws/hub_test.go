package ws_test

import (
	"testing"
	"time"

	"github.com/akaitigo/digi-engawa/api/internal/ws"
)

func TestHubJoinAndBroadcast(t *testing.T) {
	hub := ws.NewHub()
	client := ws.NewClient()
	defer client.Close()

	if err := hub.Join("room-1", client); err != nil {
		t.Fatalf("join: %v", err)
	}

	if hub.RoomCount("room-1") != 1 {
		t.Errorf("expected 1 client, got %d", hub.RoomCount("room-1"))
	}

	hub.Broadcast("room-1", ws.Message{Type: "test", Data: "hello"})

	select {
	case msg := <-client.Messages():
		if string(msg) == "" {
			t.Error("expected non-empty message")
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for broadcast message")
	}
}

func TestHubLeave(t *testing.T) {
	hub := ws.NewHub()
	client := ws.NewClient()
	defer client.Close()

	if err := hub.Join("room-1", client); err != nil {
		t.Fatalf("join: %v", err)
	}
	hub.Leave("room-1", client)

	if hub.RoomCount("room-1") != 0 {
		t.Errorf("expected 0 clients, got %d", hub.RoomCount("room-1"))
	}
}

func TestHubMultipleClients(t *testing.T) {
	hub := ws.NewHub()
	c1 := ws.NewClient()
	c2 := ws.NewClient()
	defer c1.Close()
	defer c2.Close()

	if err := hub.Join("room-1", c1); err != nil {
		t.Fatalf("join c1: %v", err)
	}
	if err := hub.Join("room-1", c2); err != nil {
		t.Fatalf("join c2: %v", err)
	}

	if hub.RoomCount("room-1") != 2 {
		t.Errorf("expected 2 clients, got %d", hub.RoomCount("room-1"))
	}

	hub.Broadcast("room-1", ws.Message{Type: "test", Data: "both"})

	for _, c := range []*ws.Client{c1, c2} {
		select {
		case msg := <-c.Messages():
			if string(msg) == "" {
				t.Error("expected non-empty message")
			}
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for broadcast")
		}
	}
}

func TestHubEmptyRoom(t *testing.T) {
	hub := ws.NewHub()
	hub.Broadcast("nonexistent", ws.Message{Type: "test", Data: "nothing"})

	if hub.RoomCount("nonexistent") != 0 {
		t.Errorf("expected 0 for empty room")
	}
}

func TestHubRoomConnectionLimit(t *testing.T) {
	hub := ws.NewHub()
	clients := make([]*ws.Client, ws.MaxClientsPerRoom)

	for i := range clients {
		clients[i] = ws.NewClient()
		defer clients[i].Close()
		if err := hub.Join("room-limit", clients[i]); err != nil {
			t.Fatalf("join client %d: %v", i, err)
		}
	}

	extra := ws.NewClient()
	defer extra.Close()
	if err := hub.Join("room-limit", extra); err == nil {
		t.Error("expected error when exceeding room limit")
	}
}
