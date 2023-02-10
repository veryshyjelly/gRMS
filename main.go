package main

import (
	"chat-app/database"
	"chat-app/modals"
	"gorm.io/gorm/logger"
)

func main() {
	database.Connect(logger.Info)

}

func GetDBSrv() modals.DBSrv {
	return &modals.DBService{
		DB: database.ChatDb,
	}
}