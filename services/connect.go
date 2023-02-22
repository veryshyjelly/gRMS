package services

import (
	"chat-app/database"
	"chat-app/services/db"
	"chat-app/services/delivery"
	"chat-app/services/msg"
)

func Connect() {
	dbservice.DBSr = dbservice.NewDBService(database.ChatDb)
	msgService.MGSr = msgService.NewMsgService(dbservice.DBSr)
	delivery.DVSr = delivery.NewDvService(msgService.MGSr)
}