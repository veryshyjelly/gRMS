package msgService

import (
	"fmt"
	"gRMS/modals"
)

// Text creates a new message with text ready to be sent to the chat
func (ms *mgs) Text(query *TextQuery) (*modals.Message, error) {
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

// Photo creates a new message with photo ready to be sent to the chat
func (ms *mgs) Photo(query *PhotoQuery) (*modals.Message, error) {
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

// Video creates a new message with video ready to be sent to the chat
func (ms *mgs) Video(query *VideoQuery) (*modals.Message, error) {
	vid, err := ms.dbs.GetVideo(query.VideoID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Video, msg.Caption = vid.(*modals.Video).ID, &query.Caption
	if query.ReplyToMessageID != 0 {
		rep, _ := ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
		msg.ReplyToMessage = rep.ID
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// Animation creates a new message with animation ready to be sent to the chat
func (ms *mgs) Animation(query *AnimationQuery) (*modals.Message, error) {
	anim, err := ms.dbs.GetAnimation(query.AnimationID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, fmt.Errorf("error creating message: %v", err)
	}

	msg.Animation, msg.Caption = anim.(*modals.Animation).ID, &query.Caption
	if query.ReplyToMessageID != 0 {
		rep, _ := ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
		msg.ReplyToMessage = rep.ID
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// Audio creates a new message with audio ready to be sent to the chat
func (ms *mgs) Audio(query *AudioQuery) (*modals.Message, error) {
	audio, err := ms.dbs.GetAudio(query.AudioID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Audio, msg.Caption = audio.(*modals.Audio).ID, &query.Caption
	if query.ReplyToMessageID != 0 {
		rep, _ := ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
		msg.ReplyToMessage = rep.ID
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// Sticker creates a new message with video ready to be sent to the chat
func (ms *mgs) Sticker(query *StickerQuery) (*modals.Message, error) {
	stk, err := ms.dbs.GetSticker(query.StickerID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Sticker = stk.(*modals.Sticker).ID
	if query.ReplyToMessageID != 0 {
		rep, _ := ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
		msg.ReplyToMessage = rep.ID
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}

// Document creates a new message with document ready to be sent to the chat
func (ms *mgs) Document(query *DocumentQuery) (*modals.Message, error) {
	doc, err := ms.dbs.GetDocument(query.DocumentID)
	if err != nil {
		return nil, err
	}

	msg, err := ms.dbs.CreateMessage(query.ChatID, query.From)
	if err != nil {
		return nil, err
	}

	msg.Document, msg.Caption = doc.(*modals.Document).ID, &query.Caption
	if query.ReplyToMessageID != 0 {
		rep, _ := ms.dbs.GetMessage(query.ReplyToMessageID, query.ChatID)
		msg.ReplyToMessage = rep.ID
	}

	err = ms.dbs.InsertMessage(msg)
	return msg, err
}