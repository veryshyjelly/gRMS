package dbservice

import (
	"chat-app/modals"
	"fmt"
)

// CreatePhoto function to create new photo entry
func (sr *DBService) CreatePhoto(filepath, filename string, thumb *modals.Photo) Media {
	photo := modals.Photo{
		Thumb: thumb,
		Metadata: &MediaMD{
			Filename: filename,
			Filepath: filepath,
		},
	}

	sr.db.Create(&photo)

	return &photo
}

// GetPhoto used to find photo by id
func (sr *DBService) GetPhoto(photoID uint64) (Media, error) {
	photo := modals.Photo{}

	sr.db.First(&photo, "id = ?", photoID)
	if photo.ID == 0 {
		return nil, fmt.Errorf("invalid photo id: %v", photoID)
	}

	return &photo, nil
}