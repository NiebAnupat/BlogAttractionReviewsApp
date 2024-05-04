package controller

import (
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/service"
	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	authService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) AuthController {
	return &AuthControllerImpl{authService: authService}
}

func (a *AuthControllerImpl) Login(c *fiber.Ctx) error {
	panic("implement me")
}

func (a *AuthControllerImpl) Logout(c *fiber.Ctx) error {
	panic("implement me")
}

func (a *AuthControllerImpl) Register(c *fiber.Ctx) error {

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := a.authService.Register(username, email, password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})

}

func (a *AuthControllerImpl) VerifyToken(c *fiber.Ctx) error {
	panic("implement me")
}

func (a *AuthControllerImpl) RefreshToken(c *fiber.Ctx) error {
	panic("unimplemented")
}
