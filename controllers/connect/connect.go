package connect

import (
	"chat-app/services"
	"chat-app/services/delivery"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ConnClient(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := services.DBS.FindUser(username, password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	websocket.New(func(c *websocket.Conn) {
		client := delivery.NewClient(user, c)
		go client.Read()
		client.Listen()
	})

	return nil
}