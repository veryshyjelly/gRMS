package modals

import (
	dbservice "chat-app/services/db"
	"time"
)

type Video struct {
	// Unique ID of the Video
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Duration of the video
	Duration uint64 `json:"duration"`
	// Thumb is the thumbnail of the video
	Thumb *Photo `json:"thumb"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
	// Metadata is the metadata of the file
	Metadata *dbservice.MediaMD
}

func (vi *Video) GetType() Filetype {
	return VideoType
}

func (vi *Video) GetFileID() uint64 {
	return vi.ID
}

func (vi *Video) GetFilesize() uint64 {
	return vi.Metadata.Filesize
}

func (vi *Video) GetFilename() string {
	return vi.Metadata.Filename
}

func (vi *Video) GetFilepath() string {
	return vi.Metadata.Filepath
}

func (vi *Video) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}