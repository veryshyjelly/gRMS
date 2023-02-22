package modals

import "time"

type ChatJoinRequest struct {
	Chat       *Chat       `json:"chat"`
	From       *User       `json:"from"`
	Date       *time.Time  `json:"date"`
	InviteLink *InviteLink `json:"invite_link"`
}
type InviteLink struct {
	InviteLink         string    `json:"invite_link"`
	Creator            *User     `json:"creator"`
	CreatesJoinRequest bool      `json:"creates_join_request"`
	IsRevoked          bool      `json:"is_revoked"`
	ExpireDate         time.Time `json:"expire_date"`
}

func NewChatInviteLink() *InviteLink {
	return &InviteLink{}
}

func NewChatJoinRequest() *ChatJoinRequest {
	return &ChatJoinRequest{}
}