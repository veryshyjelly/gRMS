package modals

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
}

type UserDB struct {
	User
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"min:6"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

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