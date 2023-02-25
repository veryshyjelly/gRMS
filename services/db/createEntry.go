package dbService

import "chat-app/modals"

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

// CreateAudio function to create a new audio entry
func (sr *DBService) CreateAudio(filepath, filename string, thumb uint64) Media {
	audio := modals.Audio{
		Filename: filename,
		Thumb:    thumb,
		Filepath: filepath,
	}

	sr.db.Create(&audio)

	return &audio
}

// CreateDocument function to create a new document entry
func (sr *DBService) CreateDocument(filepath, filename string, thumb uint64) Media {
	doc := modals.Document{
		Thumb:    thumb,
		Filename: filename,
		Filepath: filepath,
	}

	sr.db.Create(&doc)

	return &doc
}

// CreatePhoto function to create new photo entry
func (sr *DBService) CreatePhoto(filepath, filename string, thumb uint64) Media {
	photo := modals.Photo{
		Thumb:    thumb,
		Filename: filename,
		Filepath: filepath,
	}

	sr.db.Create(&photo)

	return &photo
}

// CreateSticker function to create a new sticker entry
func (sr *DBService) CreateSticker(filepath, filename, emoji string) Media {
	sticker := modals.Sticker{
		Emoji:    emoji,
		Filepath: filepath,
		Filename: filename,
	}

	sr.db.Create(&sticker)

	return &sticker
}

// CreateVideo function to create a new video entry in the database
func (sr *DBService) CreateVideo(filepath, filename string, thumb uint64) Media {
	video := modals.Video{
		Thumb:    thumb,
		Filename: filename,
		Filepath: filepath,
	}

	sr.db.Create(&video)

	return &video
}