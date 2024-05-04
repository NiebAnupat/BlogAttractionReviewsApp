package service

import (
	"io"
	"net/http"

	_AuthException "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/exception"
	_fileStorageService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/service"
	_userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
	_userService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/service"
	"github.com/google/uuid"
)

type AuthServiceImple struct {
	userService         _userService.UserService
	fileStrorageService _fileStorageService.FileStorageService
}

// Login implements AuthService.
func (a *AuthServiceImple) Login(username string, password string) (string, error) {
	panic("unimplemented")
}

// Logout implements AuthService.
func (a *AuthServiceImple) Logout(token string) error {
	panic("unimplemented")
}

// RefreshToken implements AuthService.
func (a *AuthServiceImple) RefreshToken(token string) (string, error) {
	panic("unimplemented")
}

// Register implements AuthService.
func (a *AuthServiceImple) Register(username string, email string, password string) (*_userModel.User, error) {
	ImageData, err := getDefaultAvatar()

	if err != nil {
		return nil, err
	}

	if ImageData == nil {
		return nil, &_AuthException.DownloadAvatar{}
	}

	filename := uuid.New().String() + ".svg"

	imageName, err := a.fileStrorageService.UploadFile(ImageData, filename)

	if err != nil {
		return nil, err
	}

	userCreateReq := &_userModel.UserCreateReq{
		Username: username,
		Email:    email,
		Password: password,
		Avatar:   imageName,
	}

	userModel, err := a.userService.CreateUser(userCreateReq)
	if err != nil {
		return nil, err
	}
	return userModel, nil

}

// VerifyToken implements AuthService.
func (a *AuthServiceImple) VerifyToken(token string) (string, error) {
	panic("unimplemented")
}

func getDefaultAvatar() (io.Reader, error) {

	imageURL := "https://source.boringavatars.com/beam/100/"

	// Make HTTP GET request
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, err
	}
	// defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return nil, &_AuthException.DownloadAvatar{}
	}

	imageFileReader := resp.Body

	return imageFileReader, nil
}

func NewAuthServiceImple(userService _userService.UserService, fileStrorageService _fileStorageService.FileStorageService) AuthService {
	return &AuthServiceImple{
		userService:         userService,
		fileStrorageService: fileStrorageService,
	}
}
