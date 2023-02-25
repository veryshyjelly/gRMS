package dbService

import (
	"chat-app/modals"
	"fmt"
	"time"
)

type Media interface {
	GetType() modals.Filetype
	GetFileID() uint64
	GetMetaData() *modals.MediaMD
	GetFileLinkExpiry() time.Time
}

func GetFileType(tp string) (modals.Filetype, error) {
	switch tp {
	case "photo":
		return modals.PhotoType, nil
	case "video":
		return modals.VideoType, nil
	case "audio":
		return modals.AudioType, nil
	case "document":
		return modals.DocumentType, nil
	case "sticker":
		return modals.StickerType, nil
	default:
		return 0, fmt.Errorf("invalid file typ")

	}
}

// CreateMedia is a convenience method to create a media entry
func (sr *DBService) CreateMedia(filepath, filename string, thumb uint64, filetype modals.Filetype) (Media, error) {
	switch filetype {
	case modals.PhotoType:
		return sr.CreatePhoto(filepath, filename, thumb), nil
	case modals.StickerType:
		return sr.CreateSticker(filepath, filename, ""), nil
	case modals.VideoType:
		return sr.CreateVideo(filepath, filename, thumb), nil
	case modals.AudioType:
		return sr.CreateAudio(filepath, filename, thumb), nil
	case modals.DocumentType:
		return sr.CreateDocument(filepath, filename, thumb), nil
	case modals.AnimationType:
		return sr.CreateAnimation(filepath, filename, thumb), nil
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