package ws

import (
	"sync"
)

type Client struct {
	mu   sync.Mutex
	send chan []byte
	done chan struct{}
}

func NewClient() *Client {
	return &Client{
		send: make(chan []byte, 256),
		done: make(chan struct{}),
	}
}

func (c *Client) Send(data []byte) {
	select {
	case c.send <- data:
	default:
		// Drop message if buffer is full
	}
}

func (c *Client) Messages() <-chan []byte {
	return c.send
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	select {
	case <-c.done:
		return
	default:
		close(c.done)
		close(c.send)
	}
}

func (c *Client) Done() <-chan struct{} {
	return c.done
}
