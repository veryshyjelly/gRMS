package chat

import (
	"chat-app/modals"
	"time"
)

type InviteLink struct {
	InviteLink         string       `json:"invite_link"`
	Creator            *modals.User `json:"creator"`
	CreatesJoinRequest bool         `json:"creates_join_request"`
	IsRevoked          bool         `json:"is_revoked"`
	ExpireDate         time.Time    `json:"expire_date"`
}

func NewChatInviteLink() *InviteLink {
	return &InviteLink{}
}

type JoinRequest struct {
	Chat       *modals.Chat `json:"chat"`
	From       *modals.User `json:"from"`
	Date       *time.Time   `json:"date"`
	InviteLink *InviteLink  `json:"invite_link"`
}

func NewChatJoinRequest() *JoinRequest {
	return &JoinRequest{}
}
func ApproveChatJoinRequest() {
	// TODO implement this function
}

func DeclineChatJoinRequest() {
	// TODO implement this function
}