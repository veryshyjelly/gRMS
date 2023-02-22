package dbService

import "chat-app/modals"

func (sr *DBService) AddMember(chatID uint64, userID uint64) (*modals.Participant, error) {
	chatRel := modals.Participant{
		ChatID: chatID,
		UserID: userID,
	}

	if err := sr.db.Create(&chatRel).Error; err != nil {
		return nil, err
	}

	return &chatRel, nil
}

func (sr *DBService) AddAdmin(chatId uint64, userID uint64) (*modals.Admin, error) {
	chatRel := modals.Admin{
		ChatID: chatId,
		UserID: userID,
	}

	sr.db.Create(&chatRel)

	return &chatRel, nil
}