package chat

import (
	"chat-app/modals"
	"fmt"
	"gorm.io/gorm"
)

func SetChatPhoto(db *gorm.DB, chatID uint64, photo *modals.Photo) (*modals.Chat, error) {
	chat, err := modals.GetChat(db, chatID)
	if err != nil {
		return nil, fmt.Errorf("cannot find chat: %v", err)
	}

	chat.DP = photo
	return &chat.Chat, db.Save(&chat).Error
}

func DeleteChatPhoto(db *gorm.DB, chatID uint64) (*modals.Chat, error) {
	chat, err := modals.GetChat(db, chatID)
	if err != nil {
		return nil, fmt.Errorf("cannot find chat: %v", err)
	}

	chat.DP = nil
	return &chat.Chat, db.Save(&chat).Error
}