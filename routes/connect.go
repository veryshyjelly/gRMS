package routes

import (
	"gRMS/controllers/connect"
	dbService "gRMS/services/db"
	"gRMS/services/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Connect(app *fiber.App, dbs dbService.DBS, dvs server.DVS) {
	// This function registers the connection and the signup
	// routes for the application
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", connect.ConnClient(dvs, dbs))

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.SendFile("./views/signup.html")
	})

	app.Post("/signup", connect.SignUp(dbs))
}