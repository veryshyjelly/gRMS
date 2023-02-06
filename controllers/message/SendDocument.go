package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendDocument(db *gorm.DB, query *SendDocumentQuery) (*modals.Message, error) {
	return nil, nil
}

// SendDocumentQuery is the query format for sending document
type SendDocumentQuery struct {
	From *modals.User `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// DocumentID is the file ID of the photo to be sent
	DocumentID string `json:"document"`
	// Caption is the document caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail
	Thumb *modals.Photo `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}