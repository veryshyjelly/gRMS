package modals

import "time"

type User struct {
	// Unique ID of the User
	ID uint64 `json:"id" gorm:"primaryKey"`
	// FirstName of the user
	FirstName string `json:"first_name" validate:"required,min=3"`
	// LastName of the user
	LastName string `json:"last_name"`
	// Username is unique username of the user
	Username string `json:"username" validate:"required,min=3"`
	// Bio of the user
	Bio string `json:"bio"`
	// Email is the email of the user
	Email string `json:"-" validate:"email"`
	// Password is the password of the user
	Password string `json:"-" validate:"min:6"`
	// Chats is the list of chats the user is in
	Chats []Participant `json:"chats" gorm:"foreignKey:UserID"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

type Participant struct {
	// Unique ID of the ChatId
	ID uint64 `json:"-" gorm:"primaryKey"`
	// UserID is the ID of the user
	UserID uint64 `json:"user_id"`
	// ChatID is the ID of the chat
	ChatID uint64 `json:"chat_id"`
}

type Admin struct {
	// Unique ID of the ChatId
	ID uint64 `json:"-" gorm:"primaryKey"`
	// UserID is the ID of the user
	UserID uint64 `json:"user_id"`
	// ChatID is the ID of the chat
	ChatID uint64 `json:"chat_id"`
}

func (u *User) GetUserID() uint64 {
	return u.ID
}

func (u *User) GetName() string {
	return u.FirstName + u.LastName
}

func (u *User) GetUserName() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetChats() map[uint64]bool {
	chats := make(map[uint64]bool)
	for _, chat := range u.Chats {
		chats[chat.ChatID] = true
	}
	return chats
}