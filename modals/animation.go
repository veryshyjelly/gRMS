package modals

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Animation struct {
	// Unique ID of the Animation
	ID uint64 `json:"file_id"`
	// Width of the animation
	Width uint64 `json:"width"`
	// Height of the animation
	Height uint64 `json:"height"`
	// Thumb is the thumbnail of the animation
	Thumb *Photo `json:"thumb"`
	// Duration is the time duration of the animation
	Duration time.Duration `json:"duration"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
}

type AnimationDB struct {
	Animation
	// Filepath is the path or the direct link to the file
	Filepath  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// NewAnimation to create a new animation entry
func NewAnimation(db *gorm.DB, filepath, filename string, thumb *Photo) *Animation {
	animation := AnimationDB{
		Animation: Animation{
			Filename: filename,
			Thumb:    thumb,
		},
		Filepath: filepath,
	}

	db.Create(&animation)

	return &animation.Animation
}

// FindAnimation function used to search for animation by id
func FindAnimation(db *gorm.DB, animationID uint64) (*AnimationDB, error) {
	anim := AnimationDB{Animation: Animation{ID: animationID}}
	db.First(&anim)

	if anim.Filepath == "" {
		return nil, fmt.Errorf("invalid animation id: %v", animationID)
	}

	return &anim, nil
}