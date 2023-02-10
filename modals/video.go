package modals

import (
	"fmt"
	"time"
)

type Video struct {
	// Unique ID of the Video
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Duration of the video
	Duration uint64 `json:"duration"`
	// Thumb is the thumbnail of the video
	Thumb *Photo `json:"thumb"`
	// MimeType is the mime type of the file
	MimeType string `json:"mime_type"`
	// Metadata is the metadata of the file
	Metadata *MediaMD
}

// CreateVideo function to create a new video entry in the database
func (sr *DBService) CreateVideo(filepath, filename string, thumb *Photo) *Video {
	video := Video{
		Thumb: thumb,
		Metadata: &MediaMD{
			Filename: filename,
			Filepath: filepath,
		},
	}

	sr.DB.Create(&video)

	return &video
}

// GetVideo function used to search by video by id
func (sr *DBService) GetVideo(fileID uint64) (Media, error) {
	vid := Video{}

	sr.DB.First(&vid, "id = ?", fileID)
	if vid.ID == 0 {
		return nil, fmt.Errorf("video not found")
	}

	return &vid, nil
}

func (vi *Video) GetType() Filetype {
	return VideoType
}

func (vi *Video) GetFileID() uint64 {
	return vi.ID
}

func (vi *Video) GetFilesize() uint64 {
	return vi.Metadata.Filesize
}

func (vi *Video) GetFilename() string {
	return vi.Metadata.Filename
}

func (vi *Video) GetFilepath() string {
	return vi.Metadata.Filepath
}

func (vi *Video) GetFileLinkExpiry() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}