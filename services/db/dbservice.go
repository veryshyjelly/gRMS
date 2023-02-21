package dbservice

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

type DBS interface { // DBService is the interface for all the database services
	CreateUser(firstName, lastName, username, email, password string) (*modals.User, error)
	GetUser(userID uint64) (*modals.User, error)
	FindUser(username, password string) (*modals.User, error)
	UpdateUser(user *modals.User) error
	DeleteUser(userID uint64) error
	CreateChat(users []modals.User, title string) *modals.Chat
	GetChat(chatID uint64) (*modals.Chat, error)
	UpdateChat(chat *modals.Chat) error
	SetChatPhoto(chatID uint64, photo *modals.Photo) (*modals.Chat, error)
	DeleteChatPhoto(chatID uint64) (*modals.Chat, error)
	CreatePhoto(filepath, filename string, thumb *modals.Photo) Media
	GetPhoto(photoID uint64) (Media, error)
	CreateVideo(filepath, filename string, thumb *modals.Photo) Media
	GetVideo(videoID uint64) (Media, error)
	CreateAudio(filepath, filename string, thumb *modals.Photo) Media
	GetAudio(audioID uint64) (Media, error)
	CreateDocument(filepath, filename string, thumb *modals.Photo) Media
	GetDocument(documentID uint64) (Media, error)
	CreateSticker(filepath, emoji string) Media
	GetSticker(stickerID uint64) (Media, error)
	CreateAnimation(filepath, filename string, thumb *modals.Photo) Media
	GetAnimation(animId uint64) (Media, error)
	CreateMedia(filepath, filename string, filetype modals.Filetype) (Media, error)
	GetMedia(mediaID uint64, filetype modals.Filetype) (Media, error)
	CreateMessage(chatID uint64, userID uint64) (*modals.Message, error)
	GetMessage(messageID, chatID uint64) (*modals.Message, error)
	InsertMessage(message *modals.Message) error
}

type DBService struct { // DBServiceImp is the implementation of DBService
	db *gorm.DB
}

func NewDBService(db *gorm.DB) *DBService {
	return &DBService{db: db}
}