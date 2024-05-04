package service

import (
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/repository"
	"github.com/google/uuid"

	userException "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/exception"
	_userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"

	_fileStorageService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/service"
)

type UserServiceImpl struct {
	userRepository      repository.UserRepository
	fileStrorageService _fileStorageService.FileStorageService
}

// CreateUser implements UserService.
func (u *UserServiceImpl) CreateUser(userCreateReq *model.UserCreateReq) (*_userModel.User, error) {

	uid := uuid.New().String()
	user := &entities.User{
		ID:       uid,
		Username: userCreateReq.Username,
		Password: userCreateReq.Password,
		Email:    userCreateReq.Email,
		Avatar:   userCreateReq.Avatar,
	}

	userModel, err := u.userRepository.Create(user)
	if err != nil {
		err = u.fileStrorageService.DeleteFile(userCreateReq.Avatar)
		if err != nil {
			return nil, err
		}
		return nil, &userException.UserCreate{ID: uid}
	}
	return userModel, nil
}

// GetUserByID implements UserService.
func (u *UserServiceImpl) GetUserByUsername(username string) (*model.User, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserServiceImpl(userRepository repository.UserRepository, fileStorageService _fileStorageService.FileStorageService) UserService {
	return &UserServiceImpl{
		userRepository:      userRepository,
		fileStrorageService: fileStorageService,
	}
}
