package modals

import (
	"fmt"
	"time"
)

type Photo struct {
	// Unique ID of the Photo
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Thumb is the thumbnail for the Document
	Thumb *Photo `json:"thumb"`
	// Metadata is the metadata of the file
	Metadata *MediaMD
}

// CreatePhoto function to create new photo entry
func (sr *DBService) CreatePhoto(filepath, filename string, thumb *Photo) *Photo {
	photo := Photo{
		Thumb: thumb,
		Metadata: &MediaMD{
			Filename: filename,
			Filepath: filepath,
		},
	}

	sr.DB.Create(&photo)

	return &photo
}

// GetPhoto used to find photo by id
func (sr *DBService) GetPhoto(photoID uint64) (Media, error) {
	photo := Photo{}

	sr.DB.First(&photo, "id = ?", photoID)
	if photo.ID == 0 {
		return nil, fmt.Errorf("invalid photo id: %v", photoID)
	}

	return &photo, nil
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