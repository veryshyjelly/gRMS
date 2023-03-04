package services

import (
	"gRMS/database"
	dbService "gRMS/services/db"
	msgService "gRMS/services/msg"
	"gRMS/services/server"
)

func Connect() {
	dbService.DBSr = dbService.NewDBService(database.ChatDb)
	msgService.MGSr = msgService.NewMsgService(dbService.DBSr)
	server.DVSr = server.NewDvService(msgService.MGSr, dbService.DBSr)
	go server.DVSr.Run()
}
