package delivery

import (
	"chat-app/modals"
)

type Channel struct {
	ChatID uint64
	Users  map[*Client]bool
	Mess   chan *modals.Message
	Join   chan *Client
	Leave  chan *Client
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
	for len(c.Users) > 0 {
		select {
		case client := <-c.Join:
			c.Users[client] = true
		case client := <-c.Leave:
			delete(c.Users, client)
		case msg := <-c.Mess:
			for client := range c.Users {
				client.mu.Lock()
				client.UpdateID++
				client.Mess <- modals.NewUpdate(client.UpdateID, msg)
				client.mu.Unlock()
			}
		}
	}
	delete(DVSr.Channels, c.ChatID)
}