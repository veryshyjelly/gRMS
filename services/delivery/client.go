package delivery

import (
	"chat-app/modals"
	"fmt"
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Client struct {
	ID         string
	User       *modals.User
	Chats      map[uint64]bool
	Connection *websocket.Conn
	Mess       chan *modals.Update
	UpdateID   uint64
	mu         sync.Mutex
}

func NewClient(user *modals.User, connection *websocket.Conn) *Client {
	client := &Client{
		User:       user,
		Connection: connection,
		Mess:       make(chan *modals.Update),
		Chats:      user.GetChats(),
	}

	fmt.Println("new user joined", client.User.Username)
	fmt.Println("has chats", user.Chats)

	DVSr.Users[user.ID] = client

	for chatID := range client.Chats {
		if channel, ok := DVSr.Channels[chatID]; ok {
			channel.Join <- client
		} else {
			DVSr.mu.Lock()
			channel := NewChannel(chatID)
			DVSr.Channels[chatID] = channel
			channel.Users[client] = true
			DVSr.mu.Unlock()

			go channel.Run()
		}
	}

	return client
}