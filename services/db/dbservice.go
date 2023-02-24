package dbService

import (
	"chat-app/modals"
	"gorm.io/gorm"
)

var DBSr DBS

type DBS interface { // DBService is the interface for all the database services
	CreateUser(firstName, lastName, username, email, password string) (*modals.User, error)
	GetUser(userID uint64) (*modals.User, error)
	FindUser(username string) (*modals.User, error)
	UpdateUser(user *modals.User) error
	DeleteUser(userID uint64) error
	CreateChat(users []uint64, title string) (*modals.Chat, error)
	GetChat(chatID uint64) (*modals.Chat, error)
	UpdateChat(chat *modals.Chat) error
	SetChatPhoto(chatID uint64, photo *modals.Photo) (*modals.Chat, error)
	DeleteChatPhoto(chatID uint64) (*modals.Chat, error)
	AddMember(chatID uint64, userID uint64) (*modals.Participant, error)
	AddAdmin(chatId uint64, userID uint64) (*modals.Admin, error)
	CreatePhoto(filepath, filename string, thumb uint64) Media
	GetPhoto(photoID uint64) (Media, error)
	CreateVideo(filepath, filename string, thumb uint64) Media
	GetVideo(videoID uint64) (Media, error)
	CreateAudio(filepath, filename string, thumb uint64) Media
	GetAudio(audioID uint64) (Media, error)
	CreateDocument(filepath, filename string, thumb uint64) Media
	GetDocument(documentID uint64) (Media, error)
	CreateSticker(filepath, filename, emoji string) Media
	GetSticker(stickerID uint64) (Media, error)
	CreateAnimation(filepath, filename string, thumb uint64) Media
	GetAnimation(animId uint64) (Media, error)
	CreateMedia(filepath, filename string, filetype modals.Filetype) (Media, error)
	GetMedia(mediaID uint64, filetype modals.Filetype) (Media, error)
	CreateMessage(chatID uint64, userID uint64) (*modals.Message, error)
	GetMessage(messageID, chatID uint64) (*modals.Message, error)
	GetAllMessages(chatId uint64) []*modals.Message
	InsertMessage(message *modals.Message) error
}

type DBService struct { // DBServiceImp is the implementation of DBService
	db *gorm.DB
}

func NewDBService(db *gorm.DB) *DBService {
	return &DBService{db: db}
}