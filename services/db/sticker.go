package dbservice

import (
	"chat-app/modals"
	"fmt"
)

// CreateSticker function to create a new sticker entry
func (sr *DBService) CreateSticker(filepath, emoji string) Media {
	sticker := modals.Sticker{
		Emoji: emoji,
		Metadata: &modals.MediaMD{
			Filepath: filepath,
		},
	}

	sr.db.Create(&sticker)

	return &sticker
}

// GetSticker is used to find sticker by id
func (sr *DBService) GetSticker(stickerID uint64) (Media, error) {
	sticker := modals.Sticker{}

	sr.db.First(&sticker, "id = ?", stickerID)
	if sticker.ID == 0 {
		return nil, fmt.Errorf("sticker not found")
	}

	return &sticker, nil
}