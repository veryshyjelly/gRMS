package msgService

import (
	"chat-app/modals"
)

// Text creates a new message with text ready to be sent to the chat
func (ms *MsgService) Text(query *TextQuery) (*modals.Message, error) {
	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Text = &query.Text
	if query.ReplyToMessageID != 0 {
		rep, _ := ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
		msg.ReplyToMessage = rep.ID
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// TextQuery is query format for sending message
type TextQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// Text the body of the text message
	Text string `json:"text"`
	// ReplyToMessageId is the id of the replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}