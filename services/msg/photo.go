package msgService

import (
	"chat-app/modals"
)

// Photo creates a new message with photo ready to be sent to the chat
func (ms *MsgService) Photo(query *PhotoQuery) (*modals.Message, error) {
	ph, err := ms.dbs.GetPhoto(query.PhotoID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Photo, msg.Caption = ph.(*modals.Photo).ID, &query.Caption
	if query.ReplyToMessageID != 0 {
		rep, _ := ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
		msg.ReplyToMessage = rep.ID
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// PhotoQuery is query format for sending photo
type PhotoQuery struct {
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