package chat

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

func GetChat(db *gorm.DB, chatID uint64) (*modals.Chat, error) {
	chat, err := modals.GetChat(db, chatID)
	if err != nil {
		return nil, err
	}

	return &chat.Chat, nil
}