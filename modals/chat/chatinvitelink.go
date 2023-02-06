package chatmodals

import (
	"chat-app/modals"
	"time"
)

type ChatInviteLink struct {
	InviteLink         string       `json:"invite_link"`
	Creator            *modals.User `json:"creator"`
	CreatesJoinRequest bool         `json:"creates_join_request"`
	IsRevoked          bool         `json:"is_revoked"`
	ExpireDate         time.Time    `json:"expire_date"`
}

func NewChatInviteLink() *ChatInviteLink {
	return &ChatInviteLink{}
}