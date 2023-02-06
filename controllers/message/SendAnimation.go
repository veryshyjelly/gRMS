package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendAnimation(db *gorm.DB, query *SendAnimationQuery) (*modals.Message, error) {
	return nil, nil
}

// SendAnimationQuery is query format for sending animation
type SendAnimationQuery struct {
	From *modals.User `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// AnimationID is the file ID of the photo to be sent
	AnimationID string `json:"animation"`
	// Caption is the animation caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail
	Thumb *modals.Photo `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}