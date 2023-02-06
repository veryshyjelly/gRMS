package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendVideo(db *gorm.DB, query *SendVideoQuery) (*modals.Message, error) {
	return nil, nil
}

type SendVideoQuery struct {
	From *modals.User `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// VideoID is the file ID of the photo to be sent
	VideoID string `json:"video"`
	// Caption is the video caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail
	Thumb *modals.Photo `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}