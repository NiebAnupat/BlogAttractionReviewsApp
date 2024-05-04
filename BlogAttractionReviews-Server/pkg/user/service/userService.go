package service

import (
	_userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
)

type UserService interface {
	CreateUser(userCreateReq *_userModel.UserCreateReq) (*_userModel.User, error)
	GetUserByUsername(username string) (*_userModel.User, error)
}
