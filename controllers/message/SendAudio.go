package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendAudio(db *gorm.DB, query *SendAudioQuery) (*modals.Message, error) {
	return nil, nil
}

// SendAudioQuery is query format for sending audio
type SendAudioQuery struct {
	From *modals.User `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// AudioID is the file ID of the photo to be sent
	AudioID string `json:"audio"`
	// Caption is the audio caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail
	Thumb *modals.Photo `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}