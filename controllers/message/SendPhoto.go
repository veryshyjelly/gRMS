package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendPhoto(db *gorm.DB, query *SendPhotoQuery) (*modals.Message, error) {
	ph, err := modals.FindPhoto(db, query.PhotoID)
	if err != nil {
		return nil, err
	}

	msg, err := modals.NewMessage(db, query.ChatID, query.UserID)
	if err != nil {
		return nil, err
	}

	msg.Photo, msg.Caption = &ph.Photo, &query.Caption
	msg.ReplyToMessage, err = modals.FindMessage(db, query.ReplyToMessageID, query.ChatID)

	return msg, nil
}

type SendPhotoQuery struct {
	// UserID is the id of performer of the action
	UserID uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// PhotoID is the file ID of the photo to be sent
	PhotoID uint64 `json:"photo"`
	// Caption is the photo caption
	Caption string `json:"caption"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}