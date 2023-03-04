package dbService

import (
	"fmt"
	"gRMS/modals"
)

// CreateChat creates a new chat and registers in the database
func (sr *DBService) CreateChat(users []uint64, title string) (*modals.Chat, error) {
	fmt.Printf("creating chat with users %v and title %v", users, title)

	chat := modals.Chat{
		Title:   title,
		Members: []modals.Participant{},
	}

	// First create a chat entry with the given title
	sr.db.Create(&chat)
	// Then create a table for the messages in the chat
	if err := sr.db.Table(fmt.Sprint(chat.ID)).AutoMigrate(&modals.Message{}); err != nil {
		return nil, fmt.Errorf("cannot create message table: %v for chat %v", err, chat.ID)
	}

	// Then create relation for the admin
	_, err := DBSr.AddAdmin(chat.ID, users[0])
	if err != nil {
		return nil, fmt.Errorf("error creating group: %v", err)
	}
	chat.Admins = append(chat.Admins, modals.Admin{UserID: users[0]})

	// Then create the relation for the participants in the chat
	for _, v := range users {
		// AddMember function creates the relation and returns the relation
		rel, err := DBSr.AddMember(chat.ID, v)
		if err != nil {
			return nil, fmt.Errorf("error creating group: %v", err)
		}
		chat.Members = append(chat.Members, *rel)
	}
	// return the final created chat
	return &chat, nil
}

// GetChat used to find chat using chatID
func (sr *DBService) GetChat(chatID uint64) (*modals.Chat, error) {
	chat := modals.Chat{}
	// Preload the members and admins of the chat to populate the fields
	sr.db.Preload("Members").Preload("Admins").First(&chat, "id = ?", chatID)
	if chat.ID == 0 {
		return nil, fmt.Errorf("invalid chat id %v", chatID)
	}

	return &chat, nil
}

// UpdateChat updates the chat
func (sr *DBService) UpdateChat(chat *modals.Chat) error {
	// TODO: check if chat exists
	return nil
}

// SetChatPhoto sets the chat photo
func (sr *DBService) SetChatPhoto(chatID uint64, photo *modals.Photo) (*modals.Chat, error) {
	chat, err := sr.GetChat(chatID)
	if err != nil {
		return nil, fmt.Errorf("cannot find chat: %v", err)
	}

	chat.DP = photo.ID
	return chat, sr.db.Save(&chat).Error
}

// DeleteChatPhoto deletes the chat photo
func (sr *DBService) DeleteChatPhoto(chatID uint64) (*modals.Chat, error) {
	chat, err := sr.GetChat(chatID)
	if err != nil {
		return nil, fmt.Errorf("cannot find chat: %v", err)
	}

	chat.DP = 0
	return chat, sr.db.Save(&chat).Error
}
