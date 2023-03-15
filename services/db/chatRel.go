package dbService

import "gRMS/modals"

// AddMember creates a member relation of the user with the chat
func (sr *dbs) AddMember(chatID, userID uint64) (*modals.Participant, error) {
	chatRel := modals.Participant{
		ChatID: chatID,
		UserID: userID,
	}

	if err := sr.db.Create(&chatRel).Error; err != nil {
		return nil, err
	}

	return &chatRel, nil
}

// AddAdmin creates an admin relation of the user with the chat
func (sr *dbs) AddAdmin(chatId, userID uint64) (*modals.Admin, error) {
	chatRel := modals.Admin{
		ChatID: chatId,
		UserID: userID,
	}

	sr.db.Create(&chatRel)

	return &chatRel, nil
}

// RemoveMember deletes the relation of member to the chat
func (sr *dbs) RemoveMember(chatId, userId uint64) error {
	chatRel := modals.Participant{
		ChatID: chatId,
		UserID: userId,
	}

	return sr.db.Delete(&chatRel).Error
}

// RemoveAdmin deletes the relation of user as admin
func (sr *dbs) RemoveAdmin(chatId, userID uint64) error {
	chatRel := modals.Admin{
		ChatID: chatId,
		UserID: userID,
	}

	return sr.db.Delete(&chatRel).Error
}