package server

import (
	"fmt"
	"log"
)

func (c *Client) Read() {
	defer func() {
		for chatID := range c.GetChats() {
			if channel, ok := DVSr.ActiveChannels()[chatID]; ok {
				channel.Leave <- c
			}
		}
		fmt.Println("user left", c.user.Username)
		DVSr.LeaveUser() <- c.GetUserID()
		if err := c.Connection.Close(); err != nil {
			log.Println("error while closing connection", err)
		}
	}()

	fmt.Println("reading from client", c.user.ID)

	for {
		if _, p, err := c.Connection.ReadMessage(); err != nil {
			log.Println("error while reading message from client", err)
			break
		} else {
			fmt.Println("message received from client", c.user.ID)
			DVSr.HandleReq(p, c)
		}
	}
}

func (c *Client) Listen() {
	fmt.Println("listening to client", c.user.ID)
	defer fmt.Println("stopped listening to client", c.user.ID)
	for {
		select {
		case msg := <-c.updates:
			c.UpdateID++
			msg.ID = c.UpdateID
			if err := c.Connection.WriteJSON(msg); err != nil {
				log.Println("error while writing message to client", err)
			}
		case chatID := <-c.join:
			c.mu.Lock()
			c.chats[chatID] = true
			c.mu.Unlock()
		}
	}
}