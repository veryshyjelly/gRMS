package connect

import (
	"fmt"
	"gRMS/modals"
	dbService "gRMS/services/db"
	"gRMS/services/server"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ConnClient() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		username := c.Query("username")
		password := c.Query("password")

		user, err := dbService.DBSr.FindUser(username)
		if err != nil || user.Password != password {
			err := c.WriteJSON(modals.ErrorUpdate("incorrect username or password"))
			c.Close()
			if err != nil {
				log.Println("error sending error:", err)
			}
			return
		}

		client := server.NewClient(user, c)
		fmt.Println("new client connected", client.GetUsername())
		go client.SyncHistory()
		go client.Listen()
		client.Read()
	})
}
