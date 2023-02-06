package chatmodals

import (
	"chat-app/modals"
	"time"
)

type ChatMemberUpdated struct {
	Chat *modals.Chat `json:"chat"`
	From *modals.User `json:"from"`
	Date time.Time    `json:"date"`
}

func NewChatMemberUpdated() *ChatMemberUpdated {
	return &ChatMemberUpdated{}
}