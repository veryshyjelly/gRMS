package server

import (
	"fmt"
	"gRMS/modals"
)

type Channel interface {
	GetChatID() uint64
	UserJoin() chan Client
	UserLeave() chan Client
	Message() chan *modals.Message
	Run()
}

type channel struct {
	// ChatID is the id of the chat
	ChatID uint64
	// Users is the map of all the users in the channel
	Users map[Client]bool
	// Mess is the channel to send messages to all the users in the chat
	Mess chan *modals.Message
	// Join is the chan to add a new user to the channel
	Join chan Client
	// Leave is the chan to remove a user from the channel
	Leave chan Client
}

func NewChannel(chatID uint64, user Client) Channel {
	return &channel{
		ChatID: chatID,
		Users:  map[Client]bool{user: true},
		Mess:   make(chan *modals.Message),
		Join:   make(chan Client),
		Leave:  make(chan Client),
	}
}

// Run is the main function of the channel
// that listens to the Join, Leave and Mess requests
func (c *channel) Run() {
	fmt.Println("channel started", c.ChatID)
	defer func() {
		fmt.Println("channel stopped", c.ChatID)
		DVSr.LockChannels()
		DVSr.StopChannel() <- c.ChatID
	}()

	for len(c.Users) > 0 {
		select {
		case client := <-c.Join:
			// fmt.Printf("new user %v joined in %v\n", client.GetUsername(), c.ChatID)
			c.Users[client] = true

		case client := <-c.Leave:
			// fmt.Println("user left", client.GetUsername())
			delete(c.Users, client)

		case msg := <-c.Mess:
			//fmt.Println("new message", msg)
			for client := range c.Users {
				client.Updates() <- modals.MessageUpdate(msg)
			}
		}
	}
}

// Message returns the channel to send messages to the chat
func (c *channel) Message() chan *modals.Message {
	return c.Mess
}

// GetChatID returns the chat id of the current channel
func (c *channel) GetChatID() uint64 {
	return c.ChatID
}

func (c *channel) UserJoin() chan Client {
	return c.Join
}

func (c *channel) UserLeave() chan Client {
	return c.Leave
}
