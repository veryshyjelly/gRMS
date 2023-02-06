package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendMessage(db *gorm.DB, query *SendMessageQuery) (*modals.Message, error) {
	return nil, nil
}

type SendMessageQuery struct {
	From *modals.User `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// Text the body of the text message
	Text string `json:"text"`
	// ReplyToMessageId is the id of the replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}