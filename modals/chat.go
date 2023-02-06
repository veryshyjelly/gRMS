package modals

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	ID          uint64   `json:"id" gorm:"primaryKey"`
	Title       string   `json:"title"`
	Users       []User   `json:"users"`
	Usernames   []string `json:"usernames"`
	DP          Photo    `json:"dp"`
	Description string   `json:"description"`
	InviteLink  string   `json:"invite_link"`
}

type ChatDB struct {
	Chat
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

func NewChat(db *gorm.DB, user User, title string) *Chat {
	chat := ChatDB{
		Chat: Chat{
			Title:     title,
			Users:     []User{user},
			Usernames: []string{user.Username},
		},
	}

	db.Create(&chat)
	db.Table(fmt.Sprint(chat.ID)).Create(&Message{})

	return &chat.Chat
}