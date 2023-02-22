package delivery

import (
	"chat-app/modals"
	"fmt"
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Client struct {
	ID         uint64
	UpdateID   uint64
	User       *modals.User
	Chats      map[uint64]bool
	Mess       chan *modals.Update
	Join       chan uint64
	Connection *websocket.Conn
	mu         sync.Mutex
}

func NewClient(user *modals.User, connection *websocket.Conn) *Client {
	client := &Client{
		ID:         user.ID,
		User:       user,
		Connection: connection,
		Mess:       make(chan *modals.Update),
		Join:       make(chan uint64),
		Chats:      user.GetChats(),
	}

	fmt.Println("new user joined", client.User.Username)
	fmt.Println("has chats", user.Chats)

	DVSr.mu.Lock()
	DVSr.Users[user.ID] = client
	DVSr.mu.Unlock()

	for chatID := range client.Chats {
		if channel, ok := DVSr.Channels[chatID]; ok {
			channel.Join <- client
		} else {
			DVSr.mu.Lock()
			channel := NewChannel(chatID)
			DVSr.Channels[chatID] = channel
			channel.Users[client] = true
			go channel.Run()
			DVSr.mu.Unlock()
		}
	}

	return client
}