package server

import (
	"log"
)

func (c *client) Read(dvs DVS) {
	defer func() {
		for chatID := range c.GetChats() {
			if channel, ok := dvs.ActiveChannels()[chatID]; ok {
				channel.UserLeave() <- c
			}
		}

		dvs.LockUsers()
		dvs.LeaveUser() <- c.GetUserID()
		if err := c.Connection.Close(); err != nil {
			log.Println("error while closing connection", err)
		}
	}()

	for {
		if _, p, err := c.Connection.ReadMessage(); err != nil {
			log.Println("error while reading message from client", err)
			break
		} else {
			dvs.HandleReq(c, p)
		}
	}
}

func (c *client) Listen() {
	for {
		select {
		case up := <-c.updates:
			c.UpdateID++
			up.ID = c.UpdateID
			if err := c.Connection.WriteJSON(up); err != nil {
				log.Println("error while writing message to client", err)
			}
		case his := <-c.history:
			his.ID = 0
			if err := c.Connection.WriteJSON(his); err != nil {
				log.Println("error while writing message to client", err)
			}
		case chatID := <-c.Join:
			c.Chats[chatID] = true
		}
	}
}