package server

import (
	"gRMS/modals"
	dbService "gRMS/services/db"
)

// SyncHistory function sends the chat history to the user
func (c *client) SyncHistory() {
	// Iterate on all the chats that the user has
	for chatID := range c.GetChats() {
		// Get all the messages in the chat
		messages := dbService.DBSr.GetAllMessages(chatID)
		// then concurrently send them to the user
		go c.SendAllMessages(messages)
	}
}

// SendAllMessages function sends all the messages to the user
func (c *client) SendAllMessages(mess []*modals.Message) {
	for _, m := range mess {
		c.updates <- modals.MessageUpdate(m)
	}
}
