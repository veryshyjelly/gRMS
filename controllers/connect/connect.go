package connect

import (
	"gRMS/modals"
	dbService "gRMS/services/db"
	"gRMS/services/server"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ConnClient(dvs server.DVS, dbs dbService.DBS) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		username := c.Query("username")
		password := c.Query("password")

		user, err := dbs.FindUser(username)
		if err != nil || user.Password != password {
			err := c.WriteJSON(modals.ErrorUpdate("incorrect username or password"))
			c.Close()
			if err != nil {
				log.Println("error sending error:", err)
			}
			return
		}

		client := dvs.NewClient(user, c)
		go client.SyncHistory(dbs)
		go client.Listen()
		client.Read(dvs)
	})
}