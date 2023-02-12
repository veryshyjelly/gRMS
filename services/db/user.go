package dbservice

import (
	"chat-app/modals"
	"fmt"
)

// CreateUser function to create a new user entry
func (sr *DBService) CreateUser(firstName, lastName, username, email, password string) (*modals.User, error) {
	user := modals.User{}

	sr.db.First(&user, "email = ?", email)
	if user.GetEmail() != "" {
		return nil, fmt.Errorf("email already exists")
	}

	user = modals.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Metadata: modals.UserMD{
			Email:    email,
			Password: password,
		},
	}

	sr.db.Create(&user)

	return &user, nil
}

// GetUser is used to find user by id
func (sr *DBService) GetUser(userID uint64) (*modals.User, error) {
	user := modals.User{}

	sr.db.First(&user, "id = ?", userID)
	if user.ID == 0 {
		return nil, fmt.Errorf("invalid user id: %v", userID)
	}

	return &user, nil
}
func (sr *DBService) UpdateUser(user *modals.User) error {
	return nil
}

func (sr *DBService) DeleteUser(userID uint64) error {
	return nil
}