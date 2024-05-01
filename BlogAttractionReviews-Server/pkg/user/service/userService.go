package service

import (
	_userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
)

type UserService interface {
	CreateUser(userCreateReq *_userModel.UserCreateReq) (string, error)
	GetUserByID(id string) (*_userModel.User, error)
}
