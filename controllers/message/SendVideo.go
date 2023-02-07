package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendVideo(db *gorm.DB, query *SendVideoQuery) (*modals.Message, error) {
	vid, err := modals.FindVideo(db, query.VideoID)
	if err != nil {
		return nil, err
	}

	msg, err := modals.NewMessage(db, query.ChatID, query.UserID)
	if err != nil {
		return nil, err
	}

	msg.Video, msg.Caption = &vid.Video, &query.Caption
	msg.ReplyToMessage, err = modals.FindMessage(db, query.ReplyToMessageID, query.ChatID)
	if query.Thumb != nil {
		msg.Video.Thumb = query.Thumb
	}

	return msg, nil
}

type SendVideoQuery struct {
	// UserID is the id of performer of the action
	UserID uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// VideoID is the file ID of the photo to be sent
	VideoID uint64 `json:"video"`
	// Caption is the video caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail
	Thumb *modals.Photo `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}