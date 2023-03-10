package server

import (
	"fmt"
	"gRMS/modals"
	msgService "gRMS/services/msg"
	"log"
)

type MessQuery struct {
	ChatID           uint64 `json:"chat_id"`
	Text             string `json:"text,omitempty"`
	Document         uint64 `json:"doc,omitempty"`
	Photo            uint64 `json:"photo,omitempty"`
	Audio            uint64 `json:"audio,omitempty"`
	Video            uint64 `json:"video,omitempty"`
	Animation        uint64 `json:"animation,omitempty"`
	Thumb            uint64 `json:"thumb,omitempty"`
	Caption          string `json:"caption,omitempty"`
	ReplyToMessageID uint64 `json:"reply_to_message_id,omitempty"`
}

// HandleMess handles the message query
func (dvs *DvService) HandleMess(c Client, m *MessQuery) {
	var msg *modals.Message
	var err error

	if _, ok := dvs.ActiveUsers()[c.GetUserID()].GetChats()[m.ChatID]; !ok {
		c.Updates() <- modals.ErrorUpdate("you are not a member of this chat")
		return
	}

	switch {
	case m.Text != "":
		msg, err = dvs.Mgs.Text(&msgService.TextQuery{
			From:             c.GetUserID(),
			ChatID:           m.ChatID,
			Text:             m.Text,
			ReplyToMessageID: m.ReplyToMessageID,
		})
	case m.Document != 0:
		msg, err = dvs.Mgs.Document(&msgService.DocumentQuery{
			From:             c.GetUserID(),
			ChatID:           m.ChatID,
			DocumentID:       m.Document,
			Caption:          m.Caption,
			ReplyToMessageID: m.ReplyToMessageID,
		})
	case m.Photo != 0:
		msg, err = dvs.Mgs.Photo(&msgService.PhotoQuery{
			From:             c.GetUserID(),
			ChatID:           m.ChatID,
			PhotoID:          m.Photo,
			Caption:          m.Caption,
			ReplyToMessageID: m.ReplyToMessageID,
		})
	case m.Audio != 0:
		msg, err = dvs.Mgs.Audio(&msgService.AudioQuery{
			From:             c.GetUserID(),
			ChatID:           m.ChatID,
			AudioID:          m.Audio,
			Caption:          m.Caption,
			ReplyToMessageID: m.ReplyToMessageID,
		})
	case m.Video != 0:
		msg, err = dvs.Mgs.Video(&msgService.VideoQuery{
			From:             c.GetUserID(),
			ChatID:           m.ChatID,
			VideoID:          m.Video,
			Caption:          m.Caption,
			ReplyToMessageID: m.ReplyToMessageID,
		})
	case m.Animation != 0:
		msg, err = dvs.Mgs.Animation(&msgService.AnimationQuery{
			From:             c.GetUserID(),
			ChatID:           m.ChatID,
			AnimationID:      m.Animation,
			Caption:          m.Caption,
			ReplyToMessageID: m.ReplyToMessageID,
		})
	default:
		err = fmt.Errorf("unknown message type")
	}

	if err != nil {
		e := fmt.Sprintf("error while processing message: %v", err)
		msg = &modals.Message{Text: &e}
	} else {
		dvs.SendMess(msg)
	}
}

func (dvs *DvService) SendMess(msg *modals.Message) {
	// This function sends the message where it needs to go
	if channel, ok := dvs.ActiveChannels()[msg.Chat]; ok {
		channel.Message() <- msg
	} else {
		log.Println("error channel not found")
	}
}
