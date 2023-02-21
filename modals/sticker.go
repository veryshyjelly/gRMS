package modals

import (
	dbService "chat-app/services/db"
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
	Metadata *dbService.MediaMD
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