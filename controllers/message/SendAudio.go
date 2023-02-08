package message

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func SendAudio(db *gorm.DB, query *SendAudioQuery) (*modals.Message, error) {
	audio, err := modals.FindAudio(db, query.AudioID)
	if err != nil {
		return nil, err
	}

	msg, err := modals.NewMessage(db, query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Audio, msg.Caption = &audio.Audio, &query.Caption
	msg.ReplyToMessage, err = modals.FindMessage(db, query.ReplyToMessageID, query.ChatID)
	if query.Thumb != 0 {
		thumb, err := modals.FindPhoto(db, query.Thumb)
		if err == nil {
			msg.Audio.Thumb = &thumb.Photo
		}
	}

	err = msg.Insert(db)
	return msg, err
}

// SendAudioQuery is query format for sending audio
type SendAudioQuery struct {
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