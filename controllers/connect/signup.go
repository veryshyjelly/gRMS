package connect

import (
	"chat-app/services"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	firstName := c.FormValue("firstname")
	lastName := c.FormValue("lastname")
	userName := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")

	user, err := services.DBS.CreateUser(firstName, lastName, userName, email, password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}