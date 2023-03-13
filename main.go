package main

import (
	"gRMS/database"
	"gRMS/routes"
	dbService "gRMS/services/db"
	msgService "gRMS/services/msg"
	"gRMS/services/server"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/logger"
)

func main() {
	db := database.Connect(logger.Warn)

	app := fiber.New()

	dbs := dbService.NewDBService(db)
	mgs := msgService.NewMsgService(dbs)
	dvs := server.NewDvService(mgs, dbs)
	go dvs.Run()

	routes.Connect(app, dbs, dvs)
	routes.RegMedia(app, dbs)

	log.Fatalln(app.Listen(":8080"))
}