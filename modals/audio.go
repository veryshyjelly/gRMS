package modals

import (
	"time"
)

type Audio struct {
	// Unique ID of the Audio
	ID uint64 `json:"id"`
	// Title of the audio
	Title string `json:"title"`
	// Duration is time duration of the audio
	Duration time.Duration `json:"duration"`
	// Thumb is the thumbnail for the audio
	Thumb uint64 `json:"thumb"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
	// Filename is the name of the file
	Filename string `json:"-"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"-"`
	// Filepath is the path of the file
	Filepath string `json:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

func (au Audio) GetType() Filetype {
	return AudioType
}

func (au Audio) GetFileID() uint64 {
	return au.ID
}

func (au Audio) GetMetaData() *MediaMD {
	return &MediaMD{
		Filename: au.Filename,
		Filesize: au.Filesize,
		Filepath: au.Filepath,
	}
}

func (au Audio) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}