package repository

import (
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/database"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
	_userException "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/exception"
	// "gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB database.Database
}

func NewUserRepositoryImpl(db database.Database) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(userEntity *entities.User) (*entities.User, error) {
	err := r.DB.Connect().Create(userEntity).Error
	if err != nil {
		return nil, &_userException.UserCreate{ID: userEntity.ID}
	}
	return userEntity, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*entities.User, error) {
	user := &entities.User{}

	err := r.DB.Connect().First(user, "username = ?", username).Error
	if err != nil {
		return nil, &_userException.UserNotFound{Username: username}
	}
	return user, nil
}

// FindByID implements UserRepository.
func (r *UserRepositoryImpl) FindByID(id string) (*entities.User, error) {
	user := &entities.User{}
	err := r.DB.Connect().First(user, "id = ?", id).Error
	if err != nil {
		return nil, &_userException.UserNotFound{}
	}
	return user, nil
}
