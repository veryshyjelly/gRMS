package media

import (
	dbservice "gRMS/services/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DownloadMedia(c *fiber.Ctx) error {
	fileType, err := dbservice.GetFileType(c.Params("type"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid file type",
		})
	}

	fileId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid file id",
		})
	}

	med, err := dbservice.DBSr.GetMedia(fileId, fileType)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "the request file was not found",
		})
	}

	return c.SendFile(med.GetMetaData().Filepath, false)
}
