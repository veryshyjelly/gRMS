package server

import (
	"chat-app/modals"
	"fmt"
	"sync"
)

type Channel struct {
	// ChatID is the id of the chat
	ChatID uint64
	// Users is the map of all the users in the channel
	Users map[*Client]bool
	// Mess is the channel to send messages to all the users
	Mess chan *modals.Message
	// Join is the chan to add a new user to the channel
	Join chan *Client
	// Leave is the chan to remove a user from the channel
	Leave chan *Client
	//
	mu sync.Mutex
}

func NewChannel(chatID uint64, user *Client) *Channel {
	return &Channel{
		ChatID: chatID,
		Users:  map[*Client]bool{user: true},
		Mess:   make(chan *modals.Message),
		Join:   make(chan *Client),
		Leave:  make(chan *Client),
	}
}

// Run is the main function of the channel
// that listens to the Join, Leave and Mess requests
func (c *Channel) Run() {
	fmt.Println("channel started", c.ChatID)
	for len(c.Users) > 0 {
		select {
		case client := <-c.Join:
			fmt.Printf("new user %v joined in %v\n", client.GetUsername(), c.ChatID)
			c.Users[client] = true

		case client := <-c.Leave:
			fmt.Println("user left", client.GetUsername())
			delete(c.Users, client)

		case msg := <-c.Mess:
			for client := range c.Users {
				client.Updates() <- modals.MessageUpdate(msg)
			}
		}
	}

	DVSr.StopChannel() <- c.ChatID
}