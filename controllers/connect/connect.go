package connect

import (
	dbservice "chat-app/services/db"
	"chat-app/services/delivery"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func ConnClient() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		username := c.Query("username")
		password := c.Query("password")

		user, err := dbservice.DBSr.FindUser(username)
		if err != nil || user.Password != password {
			err := c.WriteJSON(fiber.Map{"message": "Invalid username or password"})
			c.Close()
			if err != nil {
				log.Println("error sending error:", err)
			}
			return
		}

		client := delivery.NewClient(user, c)
		go client.SyncHistory()
		go client.Listen()
		client.Read()
	})
}