package services

import (
	"chat-app/database"
	"chat-app/services/db"
	"chat-app/services/delivery"
	"chat-app/services/msg"
)

func Connect() {
	dbService.DBSr = dbService.NewDBService(database.ChatDb)
	msgService.MGSr = msgService.NewMsgService(dbService.DBSr)
	delivery.DVSr = delivery.NewDvService(msgService.MGSr)
}