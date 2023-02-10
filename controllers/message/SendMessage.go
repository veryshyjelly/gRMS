package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

// SendMessage creates a new message with text ready to be sent to the chat
func SendMessage(db *gorm.DB, query *SendMessageQuery) (*modals.Message, error) {
	msg, err := modals.CreateMessage(db, query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Text = &query.Text
	msg.ReplyToMessage, _ = modals.GetMessage(db, query.ReplyToMessageID, query.ChatID)

	err = msg.Insert(db)
	return msg, err
}

// SendMessageQuery is query format for sending message
type SendMessageQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// Text the body of the text message
	Text string `json:"text"`
	// ReplyToMessageId is the id of the replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}