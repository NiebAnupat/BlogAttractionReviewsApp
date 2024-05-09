package repository

import (
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
)

type UserRepository interface {
	Create(userEntity *entities.User) (*entities.User, error)
	FindByUsername(id string) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
}
