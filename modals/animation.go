package modals

import (
	"gorm.io/gorm"
	"time"
)

type Animation struct {
	ID       uint64        `json:"file_id"`
	Width    uint64        `json:"width"`
	Height   uint64        `json:"height"`
	Thumb    *Photo        `json:"thumb"`
	Duration time.Duration `json:"duration"`
	Filename string        `json:"filename"`
	Filesize uint64        `json:"filesize"`
	MimeType string        `json:"mime_type"`
}

type AnimationDB struct {
	Animation
	Filepath  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

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