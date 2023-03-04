package server

import (
	"gRMS/modals"

	"github.com/gofiber/websocket/v2"
)

type Client interface {
	GetChats() map[uint64]bool
	GetUserID() uint64
	GetUsername() string
	ChatJoin() chan uint64
	Updates() chan *modals.Update
	SyncHistory()
	Read()
	Listen()
}

type client struct {
	UpdateID   uint64
	User       *modals.User
	Chats      map[uint64]bool
	updates    chan *modals.Update
	Join       chan uint64
	Connection *websocket.Conn
}

// NewClient function creates a new client
func NewClient(user *modals.User, connection *websocket.Conn) Client {
	client := &client{
		User:       user,
		Connection: connection,
		updates:    make(chan *modals.Update),
		Join:       make(chan uint64),
		Chats:      user.GetChats(),
	}

	DVSr.LockUsers()
	// Add the client to delivery service
	DVSr.AddUser() <- client

	// Loop through all the chats of the user
	for chatID := range client.Chats {
		// If the channel is active then add the user to active users
		if channel, ok := DVSr.ActiveChannels()[chatID]; ok {
			channel.UserJoin() <- client
		} else {
			channel := NewChannel(chatID, client)
			go channel.Run()

			DVSr.LockChannels()
			DVSr.AddChannel() <- channel
		}
	}

	return client
}

// GetChats returns the chats of the user
func (c *client) GetChats() map[uint64]bool {
	return c.Chats
}

// GetUserID returns the user id of the respective client
func (c *client) GetUserID() uint64 {
	return c.User.ID
}

// GetUsername returns the username of the client
func (c *client) GetUsername() string {
	return c.User.Username
}

func (c *client) ChatJoin() chan uint64 {
	return c.Join
}

func (c *client) Updates() chan *modals.Update {
	return c.updates
}
