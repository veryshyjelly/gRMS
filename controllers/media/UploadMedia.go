package media

import (
	dbservice "gRMS/services/db"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// UploadMedia handler handles the upload of the file,
// saves it and returns the corresponding file id to the user
func UploadMedia(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while parsing form",
		})
	}

	fileType, err := dbservice.GetFileType(c.FormValue("type"))
	if err != nil {
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

	filepath := "./database/" + c.FormValue("type") + "/" + upFile.Filename
	filename := upFile.Filename

	err = os.WriteFile(filepath, fileBytes, 0644)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while saving file",
		})
	}

	med, err := dbservice.DBSr.CreateMedia(filepath, filename, thumbID, fileType)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "file uploaded successfully",
		"media":   med,
	})
}
