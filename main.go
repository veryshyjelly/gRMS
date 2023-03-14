package main

import (
	"gRMS/database"
	"gRMS/routes"
	dbService "gRMS/services/db"
	msgService "gRMS/services/msg"
	"gRMS/services/server"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := database.Connect(logger.Warn)

	app := fiber.New(fiber.Config{BodyLimit: 21474836480})

	dbs := dbService.NewDBService(db)
	mgs := msgService.NewMsgService(dbs)
	dvs := server.NewDvService(mgs, dbs)
	go dvs.Run()

	routes.Connect(app, dbs, dvs)
	routes.RegMedia(app, dbs)

	log.Fatalln(app.Listen(":" + port))
}
