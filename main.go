package main

import (
	"chat-app/database"
	"gorm.io/gorm/logger"
)

func main() {
	database.Connect(logger.Info)
}