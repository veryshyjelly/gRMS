package modals

import (
	"time"
)

type Document struct {
	// Unique ID of the Document
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Thumb is the thumbnail for the Document
	Thumb uint64 `json:"thumb"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Filepath is the path of the file
	Filepath string `json:"filepath"`
}

func (doc Document) GetType() Filetype {
	return DocumentType
}

func (doc Document) GetFileID() uint64 {
	return doc.ID
}

func (doc Document) GetMetaData() *MediaMD {
	return &MediaMD{
		Filename: doc.Filename,
		Filesize: doc.Filesize,
		Filepath: doc.Filepath,
	}
}

func (doc Document) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}