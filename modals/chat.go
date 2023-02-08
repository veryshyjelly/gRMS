package modals

import (
	"fmt"
	"gorm.io/gorm"
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
}

type ChatDB struct {
	Chat
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

// NewChat creates a new chat and registers in the database
func NewChat(db *gorm.DB, users []User, title string) *Chat {
	chat := ChatDB{
		Chat: Chat{
			Title:     title,
			Users:     append(make([]User, 0), users...),
			Usernames: []string{},
		},
	}

	for _, v := range users {
		chat.Usernames = append(chat.Usernames, v.Username)
	}

	db.Create(&chat)

	return &chat.Chat
}

// FindChat used to find chat using chatID
func FindChat(db *gorm.DB, chatID uint64) (*ChatDB, error) {
	chat := ChatDB{Chat: Chat{ID: chatID}}
	db.First(&chat)

	if len(chat.Users) < 2 {
		return nil, fmt.Errorf("invalid chat id: %v", chatID)
	}

	return &chat, nil
}