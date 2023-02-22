package connect

import (
	dbservice "chat-app/services/db"
	"chat-app/services/delivery"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ConnClient(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := dbservice.DBSr.FindUser(username, password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	websocket.New(func(c *websocket.Conn) {
		client := delivery.NewClient(user, c)
		go client.SyncHistory()
		go client.Listen()
		client.Read()
	})

	return nil
}