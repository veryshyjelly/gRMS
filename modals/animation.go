package modals

import (
	"time"
)

type Animation struct {
	// Unique ID of the Animation
	ID uint64 `json:"file_id"`
	// Width of the Animation
	Width uint64 `json:"width"`
	// Height of the Animation
	Height uint64 `json:"height"`
	// Thumb is the thumbnail of the Animation
	Thumb uint64 `json:"thumb"`
	// Duration is the time duration of the Animation
	Duration time.Duration `json:"duration"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Filepath is the path of the file
	Filepath string `json:"filepath"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

func NewAnimation() *Animation {
	return &Animation{}
}

func (an *Animation) GetType() Filetype {
	return AnimationType
}

func (an *Animation) GetFileID() uint64 {
	return an.ID
}

func (an *Animation) GetMetaData() *MediaMD {
	return &MediaMD{
		Filename: an.Filename,
		Filesize: an.Filesize,
		Filepath: an.Filepath,
	}
}

func (an *Animation) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}