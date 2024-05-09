package controller

import (
	_fileStorageService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/service"
	"github.com/gofiber/fiber/v2"
)

type FileStorageControllerImpl struct {
	fileStorageService _fileStorageService.FileStorageService
}

// GetFile implements FileStorageController.
func (f *FileStorageControllerImpl) GetFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	file, err := f.fileStorageService.GetFile(filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get file",
			"error":   err.Error(),
		})
	}

	// file io.Reader to response
	return c.Status(fiber.StatusOK).SendStream(file)

}

func NewFileStorageController(fileStorageService _fileStorageService.FileStorageService) FileStorageController {
	return &FileStorageControllerImpl{fileStorageService: fileStorageService}
}
