package delivery

import (
	"chat-app/modals"
	"chat-app/services/db"
)

func (c *Client) SyncHistory() {
	for chatID := range c.User.GetChats() {
		messages := dbService.DBSr.GetAllMessages(chatID)
		go c.SendAllMessages(messages)
	}
}

func (c *Client) SendAllMessages(mess []*modals.Message) {
	for _, m := range mess {
		c.Mess <- modals.NewUpdate(0, m)
	}
}