package repsoitory

import "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"

type UserRepository interface {
	Create(userEntity *entities.User) (string, error)
	FindByID(id string) (*entities.User, error)
}
