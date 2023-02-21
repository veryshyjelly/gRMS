package services

import (
	"chat-app/database"
	dbservice "chat-app/services/db"
	"chat-app/services/delivery"
	msgService "chat-app/services/msg"
)

var DBS *dbservice.DBService
var MGS *msgService.MsgService
var DVS *delivery.DvService

func Connect() {
	DBS = dbservice.NewDBService(database.ChatDb)
	MGS = msgService.NewMsgService(DBS)
	DVS = delivery.NewDvService(MGS)
}