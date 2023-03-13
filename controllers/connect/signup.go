package connect

import (
	dbservice "gRMS/services/db"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SignUp(dbs dbservice.DBS) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		firstName := c.FormValue("firstname")
		lastName := c.FormValue("lastname")
		userName := c.FormValue("username")
		password := c.FormValue("password")
		email := c.FormValue("email")

		userName = strings.Join(strings.Split(userName, " "), "-")

		user, err := dbs.CreateUser(firstName, lastName, userName, email, password)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	}
}