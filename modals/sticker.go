package modals

import (
	"time"
)

type Sticker struct {
	// Unique ID of the sticker
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Emoji associated with the sticker
	Emoji string `json:"emoji"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Filepath is the path of the file
	Filepath string `json:"filepath"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

func (st Sticker) GetType() Filetype {
	return StickerType
}

func (st Sticker) GetFileID() uint64 {
	return st.ID
}

func (st Sticker) GetMetaData() *MediaMD {
	return &MediaMD{
		Filename: st.Filename,
		Filesize: st.Filesize,
		Filepath: st.Filepath,
	}
}

func (st Sticker) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}