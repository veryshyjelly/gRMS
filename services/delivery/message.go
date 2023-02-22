package delivery

import (
	"chat-app/modals"
	"chat-app/services/msg"
	"fmt"
)

type MessQuery struct {
	Text     *msgService.TextQuery     `json:"text"`
	Document *msgService.DocumentQuery `json:"doc"`
	Photo    *msgService.PhotoQuery    `json:"photo"`
	Audio    *msgService.AudioQuery    `json:"audio"`
	Video    *msgService.VideoQuery    `json:"video"`
}

func SendMessage(m *MessQuery) {
	var msg *modals.Message
	var err error

	switch {
	case m.Text != nil:
		msg, err = msgService.MGSr.Text(m.Text)
	case m.Document != nil:
		msg, err = msgService.MGSr.Document(m.Document)
	case m.Photo != nil:
		msg, err = msgService.MGSr.Photo(m.Photo)
	case m.Audio != nil:
		msg, err = msgService.MGSr.Audio(m.Audio)
	case m.Video != nil:
		msg, err = msgService.MGSr.Video(m.Video)
	default:
		err = fmt.Errorf("unknown message type")
	}

	if err != nil {
		e := fmt.Sprintf("error while processing message: %v", err)
		msg = &modals.Message{Text: &e}
	} else {
		DVSr.Send(msg)
	}
}