package dbService

import (
	"chat-app/modals"
	"fmt"
)

// CreateDocument function to create a new document entry
func (sr *DBService) CreateDocument(filepath, filename string, thumb uint64) Media {
	doc := modals.Document{
		Thumb:    thumb,
		Filename: filename,
		Filepath: filepath,
	}

	sr.db.Create(&doc)

	return &doc
}

// GetDocument used to find document by id
func (sr *DBService) GetDocument(documentID uint64) (Media, error) {
	doc := modals.Document{}

	sr.db.First(&doc, "id = ?", documentID)
	if doc.ID == 0 {
		return nil, fmt.Errorf("document not found")
	}

	return &doc, nil
}