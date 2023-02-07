package modals

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Sticker struct {
	// Unique ID of the sticker
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Emoji associated with the sticker
	Emoji string `json:"emoji"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
}

type StickerDB struct {
	Sticker
	Filepath  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// NewSticker function to create a new sticker entry
func NewSticker(db *gorm.DB, filepath, emoji string) *Sticker {
	sticker := StickerDB{
		Sticker:  Sticker{Emoji: emoji},
		Filepath: filepath,
	}

	db.Create(&sticker)

	return &sticker.Sticker
}

// FindSticker is used to find sticker by id
func FindSticker(db *gorm.DB, stickerID uint64) (*StickerDB, error) {
	sticker := StickerDB{Sticker: Sticker{ID: stickerID}}
	db.First(&sticker)

	if sticker.Filepath == "" {
		return nil, fmt.Errorf("sticker with id %v not found ", stickerID)
	}

	return &sticker, nil
}