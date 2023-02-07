package modals

import (
	"fmt"
	"gorm.io/gorm"
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
}

type UserDB struct {
	User
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"min:6"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

// NewUser function to create a new user entry
func NewUser(db *gorm.DB, firstName, lastName, username, email, password string) *User {
	user := UserDB{
		User: User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
		},
		Email:    email,
		Password: password,
	}

	db.Create(&user)

	return &user.User
}

// FindUser is used to find user by id
func FindUser(db *gorm.DB, userID uint64) (*UserDB, error) {
	user := UserDB{User: User{ID: userID}}
	db.First(&user)

	if user.Email == "" {
		return nil, fmt.Errorf("invalid user id: %v", userID)
	}

	return &user, nil
}