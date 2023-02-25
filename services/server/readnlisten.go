package server

import (
	"fmt"
	"log"
)

func (c *client) Read() {
	defer func() {
		for chatID := range c.GetChats() {
			if channel, ok := DVSr.ActiveChannels()[chatID]; ok {
				channel.UserLeave() <- c
			}
		}
		fmt.Println("user left", c.User.Username)
		DVSr.LockUsers()
		DVSr.LeaveUser() <- c.GetUserID()
		if err := c.Connection.Close(); err != nil {
			log.Println("error while closing connection", err)
		}
	}()

	fmt.Println("reading from client", c.User.ID)

	for {
		if _, p, err := c.Connection.ReadMessage(); err != nil {
			log.Println("error while reading message from client", err)
			break
		} else {
			fmt.Println("message received from client", c.User.ID)
			DVSr.HandleReq(c, p)
		}
	}
}

func (c *client) Listen() {
	fmt.Println("listening to client", c.User.ID)
	defer fmt.Println("stopped listening to client", c.User.ID)
	for {
		select {
		case msg := <-c.updates:
			c.UpdateID++
			msg.ID = c.UpdateID
			if err := c.Connection.WriteJSON(msg); err != nil {
				log.Println("error while writing message to client", err)
			}
		case chatID := <-c.Join:
			c.Chats[chatID] = true
		}
	}
}