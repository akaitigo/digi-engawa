package ws

import (
	"sync"
)

type Client struct {
	once sync.Once
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
	case <-c.done:
		return
	default:
	}

	select {
	case c.send <- data:
	default:
	}
}

func (c *Client) Messages() <-chan []byte {
	return c.send
}

func (c *Client) Close() {
	c.once.Do(func() {
		close(c.done)
		close(c.send)
	})
}

func (c *Client) Done() <-chan struct{} {
	return c.done
}
