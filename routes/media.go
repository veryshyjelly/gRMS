package routes

import (
	"gRMS/controllers/media"

	"github.com/gofiber/fiber/v2"
)

func RegMedia(app *fiber.App) {
	app.Get("/media", media.DownloadMedia)
	app.Post("/media", media.UploadMedia)
}
