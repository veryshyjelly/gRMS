package modals

import "time"

type Document struct {
	// Unique ID of the Document
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Thumb is the thumbnail for the Document
	Thumb *Photo `json:"thumb"`
}

type DocumentDB struct {
	Document
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewDocument() *Document {
	return &Document{}
}