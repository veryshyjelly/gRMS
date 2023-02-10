package message

import (
	"chat-app/modals"
	"fmt"
	"gorm.io/gorm"
)

// SendAnimation creates a new message with animation ready to be sent to the chat
func SendAnimation(db *gorm.DB, query *SendAnimationQuery) (*modals.Message, error) {
	anim, err := modals.GetAnimation(db, query.AnimationID)
	if err != nil {
		return nil, err
	}

	msg, err := modals.CreateMessage(db, query.ChatID, query.From)
	if err != nil {
		return nil, fmt.Errorf("error creating message: %v", err)
	}

	msg.Animation, msg.Caption = anim, &query.Caption
	msg.ReplyToMessage, err = modals.GetMessage(db, query.ReplyToMessageID, query.ChatID)

	if query.Thumb != 0 {
		thumb, err := modals.GetPhoto(db, query.Thumb)
		if err == nil {
			msg.Animation.Thumb = &thumb.Photo
		}
	}

	err = msg.Insert(db)
	return msg, err
}

// SendAnimationQuery is query format for sending animation
type SendAnimationQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// AnimationID is the file ID of the photo to be sent
	AnimationID uint64 `json:"animation"`
	// Caption is the animation caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail of the animation
	Thumb uint64 `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}