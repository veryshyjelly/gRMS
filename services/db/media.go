package dbservice

import (
	"chat-app/modals"
	"fmt"
	"time"
)

type Media interface {
	GetType() modals.Filetype
	GetFileID() uint64
	GetFilesize() uint64
	GetFilename() string
	GetFilepath() string
	GetFileLinkExpiry() time.Time
}

type MediaMD struct {
	// Filesize is the size of the file in kb
	Filesize uint64 `json:"filesize"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Filepath is the path of the file
	Filepath  string `json:"filepath"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

// CreateMedia is a convenience method to create a media entry
func (sr *DBService) CreateMedia(filepath, filename string, filetype modals.Filetype) (Media, error) {
	switch filetype {
	case modals.PhotoType:
		return sr.CreatePhoto(filepath, filename, nil), nil
	case modals.StickerType:
		return sr.CreateSticker(filepath, filename), nil
	case modals.VideoType:
		return sr.CreateVideo(filepath, filename, nil), nil
	case modals.AudioType:
		return sr.CreateAudio(filepath, filename, nil), nil
	case modals.DocumentType:
		return sr.CreateDocument(filepath, filename, nil), nil
	case modals.AnimationType:
		return sr.CreateAnimation(filepath, filename, nil), nil
	default:
		return nil, nil
	}
}

// GetMedia is a convenience method to get a media entry
func (sr *DBService) GetMedia(mediaID uint64, filetype modals.Filetype) (Media, error) {
	switch filetype {
	case modals.PhotoType:
		return sr.GetPhoto(mediaID)
	case modals.StickerType:
		return sr.GetSticker(mediaID)
	case modals.VideoType:
		return sr.GetVideo(mediaID)
	case modals.AudioType:
		return sr.GetAudio(mediaID)
	case modals.DocumentType:
		return sr.GetDocument(mediaID)
	case modals.AnimationType:
		return sr.GetAnimation(mediaID)
	default:
		return nil, fmt.Errorf("invalid file type")
	}
}