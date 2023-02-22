package delivery

import (
	"fmt"
	"log"
)

func (c *Client) Read() {
	defer func() {
		for chatID := range c.Chats {
			if channel, ok := DVSr.Channels[chatID]; ok {
				fmt.Println("user left", c.User.Username)
				channel.Leave <- c
			}
		}
		c.Connection.Close()
	}()

	fmt.Println("new user joined", c.User.Username)

	for {
		_, p, err := c.Connection.ReadMessage()
		if err != nil {
			return
		}
		c.HandleReq(p)
	}
}

func (c *Client) Listen() {
	fmt.Println("listening to client", c.User.ID)
	for {
		select {
		case msg := <-c.Mess:
			if err := c.Connection.WriteJSON(msg); err != nil {
				log.Println("error while writing message to client", err)
			}
		case chatID := <-c.Join:
			c.mu.Lock()
			c.Chats[chatID] = true
			c.mu.Unlock()
		}
	}
}