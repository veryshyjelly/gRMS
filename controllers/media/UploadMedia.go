package media

import (
	dbservice "chat-app/services/db"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

func UploadMedia(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while parsing form",
		})
	}

	fileType := c.FormValue("type")
	if !(fileType == "photo" || fileType == "video" || fileType == "document" || fileType == "audio" || fileType == "sticker") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid file type",
		})
	}

	upFile := form.File["file"][0]
	file, err := upFile.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while reading file",
		})
	}

	if userID, err := strconv.ParseUint(c.FormValue("userId"), 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	} else if _, err = dbservice.DBSr.GetUser(userID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	fileBytes := make([]byte, upFile.Size)
	_, err = file.Read(fileBytes)

	thumbID, err := strconv.ParseUint(c.FormValue("thumbId", "0"), 10, 64)

	filepath := "./database/" + fileType + "/" + upFile.Filename
	filename := upFile.Filename

	err = os.WriteFile(filepath, fileBytes, 0644)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while saving file",
		})
	}

	switch fileType {
	case "photo":
		ph := dbservice.DBSr.CreatePhoto(filepath, filename, thumbID)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "photo uploaded successfully",
			"photo":   ph,
		})
	case "video":
		vd := dbservice.DBSr.CreateVideo(filepath, filename, thumbID)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "video uploaded successfully",
			"video":   vd,
		})
	case "document":
		dc := dbservice.DBSr.CreateDocument(filepath, filename, thumbID)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  "document uploaded successfully",
			"document": dc,
		})
	case "audio":
		ad := dbservice.DBSr.CreateAudio(filepath, filename, thumbID)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "audio uploaded successfully",
			"audio":   ad,
		})
	case "sticker":
		st := dbservice.DBSr.CreateSticker(filepath, filename, c.FormValue("emoji"))
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "sticker uploaded successfully",
			"sticker": st,
		})
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid file type",
		})
	}
}