package modals

import (
	"time"
)

type User struct {
	// Unique ID of the User
	ID uint64 `json:"id" gorm:"primaryKey"`
	// FirstName of the user
	FirstName string `json:"first_name"`
	// LastName of the user
	LastName string `json:"last_name"`
	// Username is unique username of the user
	Username string `json:"username"`
	// Bio of the user
	Bio string `json:"bio"`
	// Metadata is the user meta data
	Metadata UserMD
}

type UserMD struct {
	Chats     map[uint64]bool `json:"chats"`
	Email     string          `json:"email" validate:"email"`
	Password  string          `json:"password" validate:"min:6"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

func NewUser() *User {
	return &User{}
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
	return u.Metadata.Email
}

func (u *User) GetPassword() string {
	return u.Metadata.Password
}

func (u *User) GetChats() map[uint64]bool {
	return u.Metadata.Chats
}