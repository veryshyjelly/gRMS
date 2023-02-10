package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

// SendPhoto creates a new message with photo ready to be sent to the chat
func SendPhoto(db *gorm.DB, query *SendPhotoQuery) (*modals.Message, error) {
	ph, err := modals.GetPhoto(db, query.PhotoID)
	if err != nil {
		return nil, err
	}

	msg, err := modals.CreateMessage(db, query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Photo, msg.Caption = &ph.Photo, &query.Caption
	msg.ReplyToMessage, err = modals.GetMessage(db, query.ReplyToMessageID, query.ChatID)
	if query.Thumb != 0 {
		thumb, err := modals.GetPhoto(db, query.Thumb)
		if err == nil {
			msg.Photo.Thumb = &thumb.Photo
		}
	}

	err = msg.Insert(db)
	return msg, err
}

// SendPhotoQuery is query format for sending photo
type SendPhotoQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// PhotoID is the file ID of the photo to be sent
	PhotoID uint64 `json:"photo"`
	// Thumb is the thumbnail of the photo
	Thumb uint64 `json:"thumb"`
	// Caption is the photo caption
	Caption string `json:"caption"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}