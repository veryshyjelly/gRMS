package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendPhoto(db *gorm.DB, query *SendPhotoQuery) (*modals.Message, error) {
	return nil, nil
}

type SendPhotoQuery struct {
	From *modals.User `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// PhotoID is the file ID of the photo to be sent
	PhotoID string `json:"photo"`
	// Caption is the photo caption
	Caption string `json:"caption"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}