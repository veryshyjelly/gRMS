package dbService

import (
	"chat-app/modals"
	"fmt"
)

// GetPhoto used to find photo by id
func (sr *DBService) GetPhoto(photoID uint64) (Media, error) {
	photo := modals.Photo{}

	sr.db.First(&photo, "id = ?", photoID)
	if photo.ID == 0 {
		return nil, fmt.Errorf("invalid photo id: %v", photoID)
	}

	return &photo, nil
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

// GetAudio is used to find audio using id
func (sr *DBService) GetAudio(audioID uint64) (Media, error) {
	audio := modals.Audio{}

	sr.db.First(&audio, "id = ?", audioID)
	if audio.ID == 0 {
		return nil, fmt.Errorf("invalid audio id: %v", audioID)
	}

	return &audio, nil
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

// GetSticker is used to find sticker by id
func (sr *DBService) GetSticker(stickerID uint64) (Media, error) {
	sticker := modals.Sticker{}

	sr.db.First(&sticker, "id = ?", stickerID)
	if sticker.ID == 0 {
		return nil, fmt.Errorf("sticker not found")
	}

	return &sticker, nil
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