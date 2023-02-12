package msgService

import (
	"chat-app/modals"
)

// Audio creates a new message with audio ready to be sent to the chat
func (ms *MsgService) Audio(query *AudioQuery) (*modals.Message, error) {
	audio, err := ms.dbs.GetAudio(query.AudioID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Audio, msg.Caption = audio.(*modals.Audio), &query.Caption
	msg.ReplyToMessage, err = ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
	if query.Thumb != 0 {
		thumb, err := ms.dbs.GetPhoto(query.Thumb)
		if err == nil {
			msg.Audio.Thumb = thumb.(*modals.Photo)
		}
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// AudioQuery is query format for sending audio
type AudioQuery struct {
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// AudioID is the file ID of the photo to be sent
	AudioID uint64 `json:"audio"`
	// Caption is the audio caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail
	Thumb uint64 `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}