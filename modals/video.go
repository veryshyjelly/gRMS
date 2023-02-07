package modals

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	// Unique ID of the Video
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Filename of the video
	Filename string `json:"filename"`
	// Duration of the video
	Duration uint64 `json:"duration"`
	// Thumb is the thumbnail of the video
	Thumb *Photo `json:"thumb"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
}

type VideoDB struct {
	Video
	Filepath  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// NewVideo function to create a new video entry
func NewVideo(db *gorm.DB, filepath, filename string, thumb *Photo) *Video {
	video := VideoDB{
		Video: Video{
			Filename: filename,
			Thumb:    thumb,
		},
		Filepath: filepath,
	}

	db.Create(&video)

	return &video.Video
}

// FindVideo function used to search by video by id
func FindVideo(db *gorm.DB, fileID uint64) (*VideoDB, error) {
	vid := VideoDB{Video: Video{ID: fileID}}
	db.First(&vid)

	if vid.Filepath == "" {
		return nil, fmt.Errorf("invalid video id: %v", fileID)
	}

	return &vid, nil
}