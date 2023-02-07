package modals

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Audio struct {
	// Unique ID of the Audio
	ID uint64 `json:"id"`
	// Title of the audio
	Title string `json:"title"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file
	Filesize uint64 `json:"filesize"`
	// Duration is time duration of the audio
	Duration time.Duration `json:"duration"`
	// Thumb is the thumbnail for the audio
	Thumb *Photo `json:"thumb"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
}

type AudioDB struct {
	Audio
	Filepath  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// NewAudio function to create a new audio entry
func NewAudio(db *gorm.DB, filepath, filename string, thumb *Photo) *Audio {
	audio := AudioDB{
		Audio: Audio{
			Filename: filename,
			Thumb:    thumb,
		},
		Filepath: filepath,
	}

	db.Create(&audio)

	return &audio.Audio
}

// FindAudio is used to find audio using id
func FindAudio(db *gorm.DB, audioID uint64) (*AudioDB, error) {
	audio := AudioDB{Audio: Audio{ID: audioID}}
	db.First(&audio)

	if audio.Filepath == "" {
		return nil, fmt.Errorf("invalid audio id: %v", audioID)
	}

	return &audio, nil
}