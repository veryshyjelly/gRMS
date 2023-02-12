package modals

import (
	"time"
)

type Chat struct {
	// Unique ID of the Chat
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Title of the Chat
	Title string `json:"title"`
	// Users is the list of users in the chat
	Users []User `json:"users"`
	// Usernames is the list of usernames in the chat
	Usernames []string `json:"usernames"`
	// DP is the display picture of the chat
	DP *Photo `json:"dp"`
	// Description is the chat description
	Description string `json:"description"`
	// InviteLink is the current active invite link
	InviteLink string `json:"invite_link"`
	// MetaData is the chat metadata
	MetaData *ChatMD
}

type ChatMD struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}