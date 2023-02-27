package services

import (
	"chat-app/database"
	"chat-app/services/db"
	"chat-app/services/msg"
	"chat-app/services/server"
)

func Connect() {
	dbService.DBSr = dbService.NewDBService(database.ChatDb)
	msgService.MGSr = msgService.NewMsgService(dbService.DBSr)
	server.DVSr = server.NewDvService(msgService.MGSr, dbService.DBSr)
	go server.DVSr.Run()
}