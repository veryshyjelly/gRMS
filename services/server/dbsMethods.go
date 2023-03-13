package server

import "gRMS/modals"

func (sr *dvs) findUser(username string) (*modals.User, error) {
	return sr.Dbs.FindUser(username)
}

func (sr *dvs) createUser(firstName string, lastName string, username string, email string, password string) (*modals.User, error) {
	return sr.Dbs.CreateUser(firstName, lastName, username, email, password)
}

func (sr *dvs) createChat(users []uint64, title string) (*modals.Chat, error) {
	return sr.Dbs.CreateChat(users, title)
}

func (sr *dvs) getChat(chatID uint64) (*modals.Chat, error) {
	return sr.Dbs.GetChat(chatID)
}

func (sr *dvs) getAllMessages(chatID uint64) []*modals.Message {
	return sr.Dbs.GetAllMessages(chatID)
}

func (sr *dvs) addMember(chatId, userId uint64) (*modals.Participant, error) {
	return sr.Dbs.AddMember(chatId, userId)
}