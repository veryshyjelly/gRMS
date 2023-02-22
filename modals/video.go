package modals

import (
	"time"
)

type Video struct {
	// Unique ID of the Video
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Duration of the video
	Duration uint64 `json:"duration"`
	// Thumb is the thumbnail of the video
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

func (vi *Video) GetType() Filetype {
	return VideoType
}

func (vi *Video) GetFileID() uint64 {
	return vi.ID
}

func (vi *Video) GetMetaData() *MediaMD {
	return &MediaMD{
		Filename: vi.Filename,
		Filesize: vi.Filesize,
		Filepath: vi.Filepath,
	}
}

func (vi *Video) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}