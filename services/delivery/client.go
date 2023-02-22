package delivery

import (
	"chat-app/modals"
	"encoding/json"
	"fmt"
	"github.com/gofiber/websocket/v2"
	"sync"

	"log"
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

	for chatID := range client.Chats {
		if channel, ok := DVSr.Channels[chatID]; ok {
			channel.Join <- client
		} else {
			channel := NewChannel(chatID)
			DVSr.Channels[chatID] = channel
			channel.Users[client] = true
			go channel.Run()
		}
	}

	return client
}

func (c *Client) Read() {
	defer func() {
		for chatID := range c.Chats {
			if channel, ok := DVSr.Channels[chatID]; ok {
				channel.Leave <- c
			}
		}
		c.Connection.Close()
	}()

	for {
		_, p, err := c.Connection.ReadMessage()
		if err != nil {
			return
		}

		req := &Req{}
		err = json.Unmarshal(p, req)
		if err != nil {
			e := fmt.Sprintf("error while unmarshaling message: %v", err)
			c.Mess <- modals.NewUpdate(0, &modals.Message{Text: &e})
		}

		if req.Message != nil {
			SendMessage(req.Message)
		}
	}
}

type Req struct {
	Message *MessQuery `json:"message"`
	//Forward *msgService.ForwardQuery  `json:"forward"`
}

func (c *Client) Listen() {
	for {
		msg := <-c.Mess
		if err := c.Connection.WriteJSON(msg); err != nil {
			log.Println("error while writing message to client", err)
		}
	}
}