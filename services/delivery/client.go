package delivery

import (
	"chat-app/modals"
	"chat-app/services"
	msgService "chat-app/services/msg"
	"encoding/json"
	"fmt"
	"github.com/gofiber/websocket/v2"

	"log"
)

type Client struct {
	ID         string
	User       *modals.User
	Chats      map[uint64]bool
	Connection *websocket.Conn
	Mess       chan *modals.Message
}

func NewClient(user *modals.User, connection *websocket.Conn) *Client {
	client := &Client{
		User:       user,
		Connection: connection,
		Mess:       make(chan *modals.Message),
		Chats:      user.GetChats(),
	}

	for chatID := range client.Chats {
		if channel, ok := services.DVS.Channels[chatID]; ok {
			channel.Join <- client
		} else {
			channel := NewChannel(chatID)
			services.DVS.Channels[chatID] = channel
			channel.Users[client] = true
			go channel.Run()
		}
	}

	return client
}

func (c *Client) Read() {
	defer func() {
		for chatID := range c.Chats {
			if channel, ok := services.DVS.Channels[chatID]; ok {
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
			c.Mess <- &modals.Message{Text: &e}
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
type MessQuery struct {
	Text     *msgService.TextQuery     `json:"text"`
	Document *msgService.DocumentQuery `json:"doc"`
	Photo    *msgService.PhotoQuery    `json:"photo"`
	Audio    *msgService.AudioQuery    `json:"audio"`
	Video    *msgService.VideoQuery    `json:"video"`
}

func SendMessage(m *MessQuery) {
	var msg *modals.Message
	var err error

	switch {
	case m.Text != nil:
		msg, err = services.MGS.Text(m.Text)
	case m.Document != nil:
		msg, err = services.MGS.Document(m.Document)
	case m.Photo != nil:
		msg, err = services.MGS.Photo(m.Photo)
	case m.Audio != nil:
		msg, err = services.MGS.Audio(m.Audio)
	case m.Video != nil:
		msg, err = services.MGS.Video(m.Video)
	default:
		err = fmt.Errorf("unknown message type")
	}

	if err != nil {
		e := fmt.Sprintf("error while processing message: %v", err)
		msg = &modals.Message{Text: &e}
	} else {
		services.DVS.Send(msg)
	}
}

func (c *Client) Listen() {
	for {
		msg := <-c.Mess
		if err := c.Connection.WriteJSON(msg); err != nil {
			log.Println("error while writing message to client", err)
		}
	}
}