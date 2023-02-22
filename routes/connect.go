package routes

import (
	"chat-app/controllers/connect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Connect(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", connect.ConnClient())

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.SendFile("./views/signup.html")
	})
	app.Post("/signup", connect.SignUp)
}