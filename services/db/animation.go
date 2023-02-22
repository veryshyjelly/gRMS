package dbService

import (
	"chat-app/modals"
	"fmt"
)

// CreateAnimation to create a new Animation entry
func (sr *DBService) CreateAnimation(filepath, filename string, thumb uint64) Media {
	animation := modals.Animation{
		Thumb:    thumb,
		Filepath: filepath,
		Filename: filename,
	}

	sr.db.Create(&animation)

	return &animation
}

// GetAnimation function used to search for Animation by id
func (sr *DBService) GetAnimation(animationID uint64) (Media, error) {
	anim := modals.Animation{}

	sr.db.First(&anim, "id = ?", animationID)
	if anim.ID == 0 {
		return nil, fmt.Errorf("invalid Animation id: %v", animationID)
	}

	return &anim, nil
}