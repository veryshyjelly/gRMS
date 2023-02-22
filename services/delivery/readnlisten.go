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
	for {
		msg := <-c.Mess
		if err := c.Connection.WriteJSON(msg); err != nil {
			log.Println("error while writing message to client", err)
		}
	}
}