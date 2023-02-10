package modals

import (
	"fmt"
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

type userDB struct {
	User
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"min:6"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

func NewUser() *User {
	return &User{}
}

// CreateUser function to create a new user entry
func (sr *DBService) CreateUser(firstName, lastName, username, email, password string) (*User, error) {
	user := userDB{}

	sr.DB.First(&user, "email = ?", email)
	if user.Email != "" {
		return nil, fmt.Errorf("email already exists")
	}

	user = userDB{
		User: User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
		},
		Email:    email,
		Password: password,
	}

	sr.DB.Create(&user)

	return &user.User, nil
}

// GetUser is used to find user by id
func (sr *DBService) GetUser(userID uint64) (*User, error) {
	user := userDB{}

	sr.DB.First(&user, "id = ?", userID)
	if user.ID == 0 {
		return nil, fmt.Errorf("invalid user id: %v", userID)
	}

	return &user.User, nil
}

func (sr *DBService) UpdateUser(user *User) error {
	return nil
}

func (sr *DBService) DeleteUser(userID uint64) error {
	return nil
}