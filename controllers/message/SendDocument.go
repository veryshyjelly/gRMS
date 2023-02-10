package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

// SendDocument creates a new message with document ready to be sent to the chat
func SendDocument(db *gorm.DB, query *SendDocumentQuery) (*modals.Message, error) {
	doc, err := modals.GetDocument(db, query.DocumentID)
	if err != nil {
		return nil, err
	}

	msg, err := modals.CreateMessage(db, query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Document, msg.Caption = &doc.Document, &query.Caption
	msg.ReplyToMessage, err = modals.GetMessage(db, query.ReplyToMessageID, query.ChatID)
	if query.Thumb != 0 {
		thumb, err := modals.GetPhoto(db, query.Thumb)
		if err == nil {
			msg.Document.Thumb = &thumb.Photo
		}
	}

	err = msg.Insert(db)
	return msg, err
}

// SendDocumentQuery is the query format for sending document
type SendDocumentQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// DocumentID is the file ID of the photo to be sent
	DocumentID uint64 `json:"document"`
	// Caption is the document caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail
	Thumb uint64 `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}