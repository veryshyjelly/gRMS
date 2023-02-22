package dbService

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
	sr.db.First(&user, "username = ?", username)
	if user.GetUserName() != "" {
		return nil, fmt.Errorf("username already exists")
	}

	user = modals.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
		Chats:     make([]modals.Participant, 0),
	}

	//validate := validator.Validate{}
	//err := validate.Struct(user)
	//if err != nil {
	//	return nil, fmt.Errorf("invalid user data: %v", err)
	//}

	sr.db.Create(&user)

	return &user, nil
}

// GetUser is used to find user by id
func (sr *DBService) GetUser(userID uint64) (*modals.User, error) {
	user := modals.User{}

	sr.db.Preload("Chats").First(&user, "id = ?", userID)
	if user.ID == 0 {
		return nil, fmt.Errorf("invalid user id: %v", userID)
	}

	return &user, nil
}

func (sr *DBService) FindUser(username string) (*modals.User, error) {
	user := modals.User{}

	sr.db.Preload("Chats").First(&user, "username = ?", username)
	if user.ID == 0 {
		return nil, fmt.Errorf("invalid username")
	}

	return &user, nil
}

func (sr *DBService) UpdateUser(user *modals.User) error {
	// TODO implement this function
	return nil
}

func (sr *DBService) DeleteUser(userID uint64) error {
	// TODO implement this function
	return nil
}