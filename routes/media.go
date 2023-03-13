package routes

import (
	"gRMS/controllers/media"
	dbService "gRMS/services/db"
	"github.com/gofiber/fiber/v2"
)

func RegMedia(app *fiber.App, dbs dbService.DBS) {
	app.Get("/media", media.DownloadMedia(dbs))
	app.Post("/media", media.UploadMedia(dbs))
}