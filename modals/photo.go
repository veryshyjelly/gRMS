package modals

import (
	"time"
)

type Photo struct {
	// Unique ID of the Photo
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Thumb is the thumbnail for the Document
	Thumb uint64 `json:"thumb"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Filepath is the path of the file
	Filepath string `json:"filepath"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

func (ph Photo) GetType() Filetype {
	return PhotoType
}

func (ph Photo) GetFileID() uint64 {
	return ph.ID
}

func (ph Photo) GetMetaData() *MediaMD {
	return &MediaMD{
		Filename: ph.Filename,
		Filesize: ph.Filesize,
		Filepath: ph.Filepath,
	}
}

func (ph Photo) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}