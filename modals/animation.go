package modals

import (
	"fmt"
	"time"
)

type Animation struct {
	// Unique ID of the Animation
	ID uint64 `json:"file_id"`
	// Width of the Animation
	Width uint64 `json:"width"`
	// Height of the Animation
	Height uint64 `json:"height"`
	// Thumb is the thumbnail of the Animation
	Thumb *Photo `json:"thumb"`
	// Duration is the time duration of the Animation
	Duration time.Duration `json:"duration"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
	// Metadata is the metadata of the file
	Metadata *MediaMD
}

func NewAnimation() *Animation {
	return &Animation{}
}

// CreateAnimation to create a new Animation entry
func (sr *DBService) CreateAnimation(filepath, filename string, thumb *Photo) *Animation {
	animation := Animation{
		Thumb: thumb,
		Metadata: &MediaMD{
			Filepath: filepath,
			Filename: filename,
		},
	}

	sr.DB.Create(&animation)

	return &animation
}

// GetAnimation function used to search for Animation by id
func (sr *DBService) GetAnimation(animationID uint64) (Media, error) {
	anim := Animation{}

	sr.DB.First(&anim, "id = ?", animationID)
	if anim.ID == 0 {
		return nil, fmt.Errorf("invalid Animation id: %v", animationID)
	}

	return &anim, nil
}

func (an *Animation) GetType() Filetype {
	return AnimationType
}

func (an *Animation) GetFileID() uint64 {
	return an.ID
}

func (an *Animation) GetFilesize() uint64 {
	return an.Metadata.Filesize
}

func (an *Animation) GetFilename() string {
	return an.Metadata.Filename
}

func (an *Animation) GetFilepath() string {
	return an.Metadata.Filepath
}

func (an *Animation) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}