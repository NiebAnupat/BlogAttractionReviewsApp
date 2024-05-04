package repository

import (
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
	userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
)

type UserRepository interface {
	Create(userEntity *entities.User) (*userModel.User, error)
	FindByUsername(id string) (*userModel.User, error)
}
