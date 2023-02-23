package server

import (
	"chat-app/modals"
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Client struct {
	ID         uint64
	UpdateID   uint64
	user       *modals.User
	chats      map[uint64]bool
	updates    chan *modals.Update
	join       chan uint64
	Connection *websocket.Conn
	mu         sync.Mutex
}

// NewClient function creates a new client
func NewClient(user *modals.User, connection *websocket.Conn) *Client {
	client := &Client{
		user:       user,
		Connection: connection,
		updates:    make(chan *modals.Update),
		join:       make(chan uint64),
		chats:      user.GetChats(),
	}

	DVSr.Lock()
	// Add the client to delivery service
	DVSr.AddUser() <- client

	// Loop through all the chats of the user
	for chatID := range client.chats {
		// If the channel is active then add the user to active users
		if channel, ok := DVSr.ActiveChannels()[chatID]; ok {
			channel.Join <- client
		} else {
			channel := NewChannel(chatID, client)
			go channel.Run()

			DVSr.Lock()
			DVSr.AddChannel() <- channel
		}
	}

	return client
}

// GetChats returns the chats of the user
func (c *Client) GetChats() map[uint64]bool {
	return c.chats
}

func (c *Client) GetUserID() uint64 {
	return c.user.ID
}

func (c *Client) GetUsername() string {
	return c.user.Username
}

func (c *Client) ChatJoin() chan uint64 {
	return c.join
}

func (c *Client) Updates() chan *modals.Update {
	return c.updates
}