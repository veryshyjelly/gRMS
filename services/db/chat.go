package dbservice

import (
	"chat-app/modals"
	"fmt"
)

// CreateChat creates a new chat and registers in the database
func (sr *DBService) CreateChat(users []modals.User, title string) *modals.Chat {
	chat := modals.Chat{
		Title:     title,
		Users:     append(make([]modals.User, 0), users...),
		Usernames: []string{},
	}

	for _, v := range users {
		chat.Usernames = append(chat.Usernames, v.GetUserName())
	}

	sr.db.Create(&chat)

	return &chat
}

// GetChat used to find chat using chatID
func (sr *DBService) GetChat(chatID uint64) (*modals.Chat, error) {
	chat := modals.Chat{}

	sr.db.First(&chat, "id = ?", chatID)
	if chat.ID == 0 {
		return nil, fmt.Errorf("invalid chat id %v", chatID)
	}

	return &chat, nil
}

func (sr *DBService) UpdateChat(chat *modals.Chat) error {
	return nil
}