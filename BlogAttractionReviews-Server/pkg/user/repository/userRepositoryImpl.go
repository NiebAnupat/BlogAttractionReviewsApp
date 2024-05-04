package repository

import (
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/database"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
	_userException "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/exception"
	userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
	// "gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB database.Database
}

func NewUserRepositoryImpl(db database.Database) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(userEntity *entities.User) (*userModel.User, error) {
	err := r.DB.Connect().Create(userEntity).Error
	if err != nil {
		return nil, &_userException.UserCreate{ID: userEntity.ID}
	}
	return userEntity.ToUserModel(), nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*userModel.User, error) {
	user := &entities.User{}
	err := r.DB.Connect().First(user, username).Error
	if err != nil {
		return nil, err
	}
	return user.ToUserModel(), nil
}
