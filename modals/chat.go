package modals

import (
	"fmt"
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

// CreateChat creates a new chat and registers in the database
func (sr *DBService) CreateChat(users []User, title string) *Chat {
	chat := Chat{
		Title:     title,
		Users:     append(make([]User, 0), users...),
		Usernames: []string{},
	}

	for _, v := range users {
		chat.Usernames = append(chat.Usernames, v.Username)
	}

	sr.DB.Create(&chat)

	return &chat
}

// GetChat used to find chat using chatID
func (sr *DBService) GetChat(chatID uint64) (*Chat, error) {
	chat := Chat{}

	sr.DB.First(&chat, "id = ?", chatID)
	if chat.ID == 0 {
		return nil, fmt.Errorf("invalid chat id %v", chatID)
	}

	return &chat, nil
}

func (sr *DBService) UpdateChat(chat *Chat) error {
	return nil
}