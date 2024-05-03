package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	VerifyToken(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}
