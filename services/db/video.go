package dbservice

import (
	"chat-app/modals"
	"fmt"
)

// CreateVideo function to create a new video entry in the database
func (sr *DBService) CreateVideo(filepath, filename string, thumb *modals.Photo) Media {
	video := modals.Video{
		Thumb: thumb,
		Metadata: &MediaMD{
			Filename: filename,
			Filepath: filepath,
		},
	}

	sr.db.Create(&video)

	return &video
}

// GetVideo function used to search by video by id
func (sr *DBService) GetVideo(fileID uint64) (Media, error) {
	vid := modals.Video{}

	sr.db.First(&vid, "id = ?", fileID)
	if vid.ID == 0 {
		return nil, fmt.Errorf("video not found")
	}

	return &vid, nil
}