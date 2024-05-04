package controller

import (
	"errors"

	_authException "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/exception"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/service"
	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	authService service.AuthService
}

func (a *AuthControllerImpl) Login(c *fiber.Ctx) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := a.authService.Login(username, password)
	if err != nil {
		if errors.Is(err, &_authException.WrongPassword{}) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Wrong password",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successfully",
		"token":   token,
	})

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

	token := c.FormValue("token")

	username, err := a.authService.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Token is valid",
		"username": username,
	})
}

func (a *AuthControllerImpl) RefreshToken(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewAuthControllerImpl(authService service.AuthService) AuthController {
	return &AuthControllerImpl{authService: authService}
}
