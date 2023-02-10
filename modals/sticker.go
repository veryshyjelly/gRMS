package modals

import (
	"fmt"
	"time"
)

type Sticker struct {
	// Unique ID of the sticker
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Emoji associated with the sticker
	Emoji string `json:"emoji"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Metadata is the metadata of the file
	Metadata *MediaMD
}

// CreateSticker function to create a new sticker entry
func (sr *DBService) CreateSticker(filepath, emoji string) *Sticker {
	sticker := Sticker{
		Emoji: emoji,
		Metadata: &MediaMD{
			Filepath: filepath,
		},
	}

	sr.DB.Create(&sticker)

	return &sticker
}

// GetSticker is used to find sticker by id
func (sr *DBService) GetSticker(stickerID uint64) (Media, error) {
	sticker := Sticker{}

	sr.DB.First(&sticker, "id = ?", stickerID)
	if sticker.ID == 0 {
		return nil, fmt.Errorf("sticker not found")
	}

	return &sticker, nil
}

func (st Sticker) GetType() Filetype {
	return StickerType
}

func (st Sticker) GetFileID() uint64 {
	return st.ID
}

func (st Sticker) GetFilesize() uint64 {
	return st.Filesize
}

func (st Sticker) GetFilename() string {
	return st.Metadata.Filename
}

func (st Sticker) GetFilepath() string {
	return st.Metadata.Filepath
}

func (st Sticker) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}