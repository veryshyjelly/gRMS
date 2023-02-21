package routes

import (
	"chat-app/controllers/connect"
	"github.com/gofiber/fiber/v2"
)

func Connect(app *fiber.App) {
	app.Get("/ws", connect.ConnClient)

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.SendFile("./views/signup.html")
	})
	app.Post("/signup", connect.SignUp)
}