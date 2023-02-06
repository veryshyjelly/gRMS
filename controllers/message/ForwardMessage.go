package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func ForwardMessage(db *gorm.DB, query *ForwardMessageQuery) (*modals.Message, error) {
	return nil, nil
}

// ForwardMessageQuery is the query format for forwarding message
type ForwardMessageQuery struct {
	From *modals.User `json:"from"`
	// ChatID is the id of the target chat
	ChatID uint64 `json:"chat_id"`
	// FromChatID is the id of the original chat
	FromChatID uint64 `json:"from_chat_id"`
	// MessageID is the id of the original message
	MessageID uint64 `json:"message_id"`
}