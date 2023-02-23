package delivery

import (
	"chat-app/modals"
	"chat-app/services/db"
)

// SyncHistory function sends the chat history to the user
func (c *Client) SyncHistory() {
	// Iterate on all the chats that the user has
	for chatID := range c.user.GetChats() {
		// Get all the messages in the chat
		messages := dbService.DBSr.GetAllMessages(chatID)
		// then concurrently send them to the user
		go c.SendAllMessages(messages)
	}
}

// SendAllMessages function sends all the messages to the user
func (c *Client) SendAllMessages(mess []*modals.Message) {
	for _, m := range mess {
		c.updates <- modals.MessageUpdate(m)
	}
}