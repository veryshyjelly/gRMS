package main

import (
	"gRMS/database"
	"gRMS/routes"
	"gRMS/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/logger"
)

func main() {
	// Basic philosophy of this project is to
	// run server at any place
	// and then connect to it using the ip address
	// and the frontend will be hosted somewhere else
	database.Connect(logger.Info)
	services.Connect()

	app := fiber.New()
	routes.Connect(app)
	routes.RegMedia(app)

	log.Fatalln(app.Listen(":8080"))
}
