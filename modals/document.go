package modals

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

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
	Filepath  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// NewDocument function to create a new document entry
func NewDocument(db *gorm.DB, filepath, filename string, thumb *Photo) *Document {
	doc := DocumentDB{
		Document: Document{
			Filename: filename,
			Thumb:    thumb,
		},
		Filepath: filepath,
	}

	db.Create(&doc)

	return &doc.Document
}

// FindDocument used to find document by id
func FindDocument(db *gorm.DB, documentID uint64) (*DocumentDB, error) {
	doc := DocumentDB{Document: Document{ID: documentID}}
	db.First(&doc)

	if doc.Filepath == "" {
		return nil, fmt.Errorf("requested data not found")
	}

	return &doc, nil
}