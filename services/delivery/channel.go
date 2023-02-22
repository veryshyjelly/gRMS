package delivery

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

func NewChannel(chatID uint64) *Channel {
	return &Channel{
		ChatID: chatID,
		Users:  make(map[*Client]bool),
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
			fmt.Printf("new user %v joined in %v\n", client.User.Username, c.ChatID)
			c.mu.Lock()
			c.Users[client] = true
			c.mu.Unlock()

		case client := <-c.Leave:
			fmt.Println("user left", client.User.Username)
			c.mu.Lock()
			delete(c.Users, client)
			c.mu.Unlock()

		case msg := <-c.Mess:
			for client := range c.Users {
				client.mu.Lock()
				client.UpdateID++
				client.Mess <- modals.NewUpdate(client.UpdateID, msg)
				client.mu.Unlock()
			}
		}
	}
	fmt.Println("channel closed", c.ChatID)
	delete(DVSr.Channels, c.ChatID)
}