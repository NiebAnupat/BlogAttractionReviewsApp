package service

import (
	_userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Logout(token string) error

	VerifyToken(token string) (*_userModel.User, error)

	RefreshToken(token string) (string, error)

	Register(username, email, password string) (*_userModel.User, error)
}
