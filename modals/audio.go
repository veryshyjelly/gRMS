package modals

import (
	"fmt"
	"time"
)

type Audio struct {
	// Unique ID of the Audio
	ID uint64 `json:"id"`
	// Title of the audio
	Title string `json:"title"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Duration is time duration of the audio
	Duration time.Duration `json:"duration"`
	// Thumb is the thumbnail for the audio
	Thumb *Photo `json:"thumb"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
	// Metadata is the metadata of the file
	Metadata *MediaMD
}

// CreateAudio function to create a new audio entry
func (sr *DBService) CreateAudio(filepath, filename string, thumb *Photo) *Audio {
	audio := Audio{
		Filename: filename,
		Thumb:    thumb,
		Metadata: &MediaMD{
			Filepath: filepath,
		},
	}

	sr.DB.Create(&audio)

	return &audio
}

// GetAudio is used to find audio using id
func (sr *DBService) GetAudio(audioID uint64) (Media, error) {
	audio := Audio{}

	sr.DB.First(&audio, "id = ?", audioID)
	if audio.ID == 0 {
		return nil, fmt.Errorf("invalid audio id: %v", audioID)
	}

	return &audio, nil
}

func (au Audio) GetType() Filetype {
	return AudioType
}

func (au Audio) GetFileID() uint64 {
	return au.ID
}

func (au Audio) GetFilesize() uint64 {
	return au.Metadata.Filesize
}

func (au Audio) GetFilename() string {
	return au.Filename
}

func (au Audio) GetFilepath() string {
	return au.Metadata.Filepath
}

func (au Audio) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}