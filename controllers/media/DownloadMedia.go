package media

import (
	dbservice "chat-app/services/db"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func DownloadMedia(c *fiber.Ctx) error {
	fileType := c.Params("type")
	fileId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid file id",
		})
	}

	switch fileType {
	case "photo":
		ph, err := dbservice.DBSr.GetPhoto(fileId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid file id",
			})
		}
		return c.SendFile(ph.GetMetaData().Filepath, false)
	case "video":
		vd, err := dbservice.DBSr.GetVideo(uint64(fileId))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid file id",
			})
		}
		return c.SendFile(vd.GetMetaData().Filepath, false)
	case "document":
		dc, err := dbservice.DBSr.GetDocument(uint64(fileId))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid file id",
			})
		}
		return c.SendFile(dc.GetMetaData().Filepath, false)
	case "audio":
		ad, err := dbservice.DBSr.GetAudio(uint64(fileId))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid file id",
			})
		}
		return c.SendFile(ad.GetMetaData().Filepath, false)
	case "sticker":
		st, err := dbservice.DBSr.GetSticker(uint64(fileId))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid file id",
			})
		}
		return c.SendFile(st.GetMetaData().Filepath, false)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid file type",
		})
	}
}