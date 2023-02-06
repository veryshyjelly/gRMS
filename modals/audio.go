package modals

import "time"

type Audio struct {
	ID       uint64        `json:"id"`
	Title    string        `json:"title"`
	Filename string        `json:"filename"`
	Filesize uint64        `json:"filesize"`
	Duration time.Duration `json:"duration"`
	Thumb    *Photo        `json:"thumb"`
	MimeType string        `json:"mime_type"`
}

type AudioDB struct {
	Audio
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewAudio() *Audio {
	return &Audio{}
}