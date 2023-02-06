package modals

import "time"

type Video struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Duration uint64 `json:"duration"`
	Title    string `json:"title"`
	Filename string `json:"filename"`
	Thumb    *Photo `json:"thumb"`
	MimeType string `json:"mime_type"`
}

type VideoDB struct {
	Video
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewVideo() *Video {
	return &Video{}
}