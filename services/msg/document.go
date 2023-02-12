package msgService

import (
	"chat-app/modals"
)

// Document creates a new message with document ready to be sent to the chat
func (ms *MsgService) Document(query *DocumentQuery) (*modals.Message, error) {
	doc, err := ms.dbs.GetDocument(query.DocumentID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Document, msg.Caption = doc.(*modals.Document), &query.Caption
	msg.ReplyToMessage, err = ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
	if query.Thumb != 0 {
		thumb, err := ms.dbs.GetPhoto(query.Thumb)
		if err == nil {
			msg.Document.Thumb = thumb.(*modals.Photo)
		}
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// DocumentQuery is the query format for sending document
type DocumentQuery struct {
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