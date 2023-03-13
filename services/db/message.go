package dbService

import (
	"fmt"
	"gRMS/modals"
)

// CreateMessage creates a new message populated with Chat and user
func (sr *dbs) CreateMessage(chatID, userID uint64) (*modals.Message, error) {
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
func (sr *dbs) GetMessage(messageID, chatID uint64) (*modals.Message, error) {
	mess := modals.Message{}

	sr.db.Table(fmt.Sprint(chatID)).First(&mess, "id = ?", messageID)
	if mess.ID == 0 {
		return nil, fmt.Errorf("invalid message id %v or chat id %v", messageID, chatID)
	}

	return &mess, nil
}

func (sr *dbs) GetAllMessages(chatID uint64) []*modals.Message {
	var mess []*modals.Message

	sr.db.Table(fmt.Sprint(chatID)).Find(&mess)

	return mess
}

func (sr *dbs) InsertMessage(m *modals.Message) error {
	return sr.db.Table(fmt.Sprint(m.Chat)).Create(m).Error
}