package dbService

import (
	"chat-app/modals"
	"fmt"
)

// CreateMessage creates a new message populated with Chat and User
func (sr *DBService) CreateMessage(chatID, userID uint64) (*modals.Message, error) {
	chat, err := sr.GetChat(chatID)
	if err != nil {
		return nil, err
	}

	user, err := sr.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return &modals.Message{Chat: chat.ID, From: user.ID}, nil
}

// GetMessage used to find message in the chat table
func (sr *DBService) GetMessage(messageID, chatID uint64) (*modals.Message, error) {
	mess := modals.Message{}

	sr.db.Table(fmt.Sprint(chatID)).First(&mess, "id = ?", messageID)
	if mess.ID == 0 {
		return nil, fmt.Errorf("invalid message id %v or chat id %v", messageID, chatID)
	}

	return &mess, nil
}

func (sr *DBService) GetAllMessages(chatID uint64) []*modals.Message {
	var mess []*modals.Message

	sr.db.Table(fmt.Sprint(chatID)).Find(&mess)

	return mess
}

func (sr *DBService) InsertMessage(m *modals.Message) error {
	return sr.db.Table(fmt.Sprint(m.Chat)).Create(m).Error
}