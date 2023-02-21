package modals

import (
	dbService "chat-app/services/db"
	"time"
)

type Document struct {
	// Unique ID of the Document
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Thumb is the thumbnail for the Document
	Thumb *Photo `json:"thumb"`
	// Metadata is the metadata of the file
	Metadata *dbService.MediaMD
}

func (doc Document) GetType() Filetype {
	return DocumentType
}

func (doc Document) GetFileID() uint64 {
	return doc.ID
}

func (doc Document) GetFilesize() uint64 {
	return doc.Metadata.Filesize
}

func (doc Document) GetFilename() string {
	return doc.Metadata.Filename
}

func (doc Document) GetFilepath() string {
	return doc.Metadata.Filepath
}

func (doc Document) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}