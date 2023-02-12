package modals

import (
	"time"
)

type Photo struct {
	// Unique ID of the Photo
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Thumb is the thumbnail for the Document
	Thumb *Photo `json:"thumb"`
	// Metadata is the metadata of the file
	Metadata *db.MediaMD
}

func (ph Photo) GetType() Filetype {
	return PhotoType
}

func (ph Photo) GetFileID() uint64 {
	return ph.ID
}

func (ph Photo) GetFilesize() uint64 {
	return ph.Metadata.Filesize
}

func (ph Photo) GetFilename() string {
	return ph.Metadata.Filename
}

func (ph Photo) GetFilepath() string {
	return ph.Metadata.Filepath
}

func (ph Photo) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}