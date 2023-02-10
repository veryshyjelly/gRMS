package modals

import "gorm.io/gorm"

type DBSrv interface { // DBService is the interface for all the database services
	CreateUser(firstName, lastName, username, email, password string) (*User, error)
	GetUser(userID uint64) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(userID uint64) error
	CreateChat(users []User, title string) *Chat
	GetChat(chatID uint64) (*Chat, error)
	UpdateChat(chat *Chat) error
	CreatePhoto(filepath, filename string, thumb *Photo) *Photo
	GetPhoto(photoID uint64) (Media, error)
	CreateVideo(filepath, filename string, thumb *Photo) *Video
	GetVideo(videoID uint64) (Media, error)
	CreateAudio(filepath, filename string, thumb *Photo) *Audio
	GetAudio(audioID uint64) (Media, error)
	CreateDocument(filepath, filename string, thumb *Photo) *Document
	GetDocument(documentID uint64) (Media, error)
	CreateSticker(filepath, emoji string) *Sticker
	GetSticker(stickerID uint64) (Media, error)
	CreateAnimation(filepath, filename string, thumb *Photo) *Animation
	GetAnimation(animId uint64) (Media, error)
	CreateMedia(filepath, filename string, filetype Filetype) (Media, error)
	GetMedia(mediaID uint64, filetype Filetype) (Media, error)
}

type DBService struct { // DBServiceImp is the implementation of DBService
	DB *gorm.DB
}