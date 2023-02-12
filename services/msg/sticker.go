package msgService

import "chat-app/modals"

// Sticker creates a new message with video ready to be sent to the chat
func (ms *MsgService) Sticker(query *StickerQuery) (*modals.Message, error) {
	stk, err := ms.dbs.GetSticker(query.StickerID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Sticker = stk.(*modals.Sticker)
	msg.ReplyToMessage, err = ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// StickerQuery is query format for sending video
type StickerQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// VideoID is the file ID of the video to be sent
	StickerID uint64 `json:"video"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}