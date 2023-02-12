package dbservice

import (
	"chat-app/modals"
	"fmt"
)

// CreateAudio function to create a new audio entry
func (sr *DBService) CreateAudio(filepath, filename string, thumb *modals.Photo) Media {
	audio := modals.Audio{
		Filename: filename,
		Thumb:    thumb,
		Metadata: &MediaMD{
			Filepath: filepath,
		},
	}

	sr.db.Create(&audio)

	return &audio
}

// GetAudio is used to find audio using id
func (sr *DBService) GetAudio(audioID uint64) (Media, error) {
	audio := modals.Audio{}

	sr.db.First(&audio, "id = ?", audioID)
	if audio.ID == 0 {
		return nil, fmt.Errorf("invalid audio id: %v", audioID)
	}

	return &audio, nil
}