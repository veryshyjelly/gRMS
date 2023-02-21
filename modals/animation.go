package modals

import (
	dbService "chat-app/services/db"
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
	Thumb *Photo `json:"thumb"`
	// Duration is the time duration of the Animation
	Duration time.Duration `json:"duration"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
	// Metadata is the metadata of the file
	Metadata *dbService.MediaMD
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

func (an *Animation) GetFilesize() uint64 {
	return an.Metadata.Filesize
}

func (an *Animation) GetFilename() string {
	return an.Metadata.Filename
}

func (an *Animation) GetFilepath() string {
	return an.Metadata.Filepath
}

func (an *Animation) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}