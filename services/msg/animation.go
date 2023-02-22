package msgService

import (
	"chat-app/modals"
	"fmt"
)

// Animation creates a new message with animation ready to be sent to the chat
func (ms *MsgService) Animation(query *AnimationQuery) (*modals.Message, error) {
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

// AnimationQuery is query format for sending animation
type AnimationQuery struct {
	// From is the user who sent the message
	From uint64 `json:"from"`
	// ChatID is the ID of the target chat
	ChatID uint64 `json:"chat_id"`
	// AnimationID is the file ID of the photo to be sent
	AnimationID uint64 `json:"animation"`
	// Caption is the animation caption
	Caption string `json:"caption"`
	// Thumb is the thumbnail of the animation
	Thumb uint64 `json:"thumb"`
	// ReplyToMessageID is the id of replied message
	ReplyToMessageID uint64 `json:"reply_to_message_id"`
}