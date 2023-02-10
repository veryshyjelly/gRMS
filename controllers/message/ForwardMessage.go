package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

// ForwardMessage forwards a message to a chat from another chat
func ForwardMessage(db *gorm.DB, query *ForwardMessageQuery) (*modals.Message, error) {
	fr, err := modals.GetMessage(db, query.MessageID, query.FromChatID)
	if err != nil {
		return nil, err
	}

	msg, err := modals.CreateMessage(db, query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Text = fr.Text
	msg.Photo = fr.Photo
	msg.Audio = fr.Audio
	msg.Document = fr.Document
	msg.Video = fr.Video
	msg.Caption = fr.Caption

	err = msg.Insert(db)
	return msg, err
}

// ForwardMessageQuery is the query format for forwarding message
type ForwardMessageQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the id of the target chat
	ChatID uint64 `json:"chat_id"`
	// FromChatID is the id of the original chat
	FromChatID uint64 `json:"from_chat_id"`
	// MessageID is the id of the original message
	MessageID uint64 `json:"message_id"`
}