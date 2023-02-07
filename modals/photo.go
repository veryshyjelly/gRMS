package modals

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	// Unique ID of the Photo
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Thumb is the thumbnail for the Document
	Thumb *Photo `json:"thumb"`
}

type PhotoDB struct {
	Photo
	Filepath  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

// NewPhoto function to create new photo entry
func NewPhoto(db *gorm.DB, filepath, filename string, thumb *Photo) *Photo {
	photo := PhotoDB{
		Photo: Photo{
			Filename: filename,
			Thumb:    thumb,
		},
		Filepath: filepath,
	}

	db.Create(&photo)

	return &photo.Photo
}

// FindPhoto used to find photo by id
func FindPhoto(db *gorm.DB, photoID uint64) (*PhotoDB, error) {
	photo := PhotoDB{Photo: Photo{ID: photoID}}
	db.First(&photo)

	if photo.Filepath == "" {
		return nil, fmt.Errorf("photo not found")
	}

	return &photo, nil
}