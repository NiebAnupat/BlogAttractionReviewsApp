package controller

import "github.com/gofiber/fiber/v2"

type FileStorageController interface {
	GetFile(c *fiber.Ctx) error
	// DeleteFile(c *fiber.Ctx) error
}
