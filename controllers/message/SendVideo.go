package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

// SendVideo creates a new message with video ready to be sent to the chat
func SendVideo(db *gorm.DB, query *SendVideoQuery) (*modals.Message, error) {
	vid, err := modals.GetVideo(db, query.VideoID)
	if err != nil {
		return nil, err
	}

	msg, err := modals.CreateMessage(db, query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Video, msg.Caption = &vid.Video, &query.Caption
	msg.ReplyToMessage, err = modals.GetMessage(db, query.ReplyToMessageID, query.ChatID)
	if query.Thumb != nil {
		msg.Video.Thumb = query.Thumb
	}

	return msg, nil
}

// SendVideoQuery is query format for sending video
type SendVideoQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// VideoID is the file ID of the video to be sent
	VideoID uint64 `json:"video"`
	// Caption is the video caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail of the video
	Thumb *modals.Photo `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}