package modals

import (
	"fmt"
	"time"
)

type Document struct {
	// Unique ID of the Document
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Thumb is the thumbnail for the Document
	Thumb *Photo `json:"thumb"`
	// Metadata is the metadata of the file
	Metadata *MediaMD
}

// CreateDocument function to create a new document entry
func (sr *DBService) CreateDocument(filepath, filename string, thumb *Photo) *Document {
	doc := Document{
		Thumb: thumb,
		Metadata: &MediaMD{
			Filename: filename,
			Filepath: filepath,
		},
	}

	sr.DB.Create(&doc)

	return &doc
}

// GetDocument used to find document by id
func (sr *DBService) GetDocument(documentID uint64) (Media, error) {
	doc := Document{}

	sr.DB.First(&doc, "id = ?", documentID)
	if doc.ID == 0 {
		return nil, fmt.Errorf("document not found")
	}

	return &doc, nil
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