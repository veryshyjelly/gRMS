package chatmodals

import (
	"chat-app/modals"
	"time"
)

type ChatJoinRequest struct {
	Chat       *modals.Chat    `json:"chat"`
	From       *modals.User    `json:"from"`
	Date       *time.Time      `json:"date"`
	InviteLink *ChatInviteLink `json:"invite_link"`
}

func NewChatJoinRequest() *ChatJoinRequest {
	return &ChatJoinRequest{}
}