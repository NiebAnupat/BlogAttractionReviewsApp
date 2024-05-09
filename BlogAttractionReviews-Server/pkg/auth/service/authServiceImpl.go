package service

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/config"
	_AuthException "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/exception"
	_fileStorageService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/service"
	_userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
	_userRepository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/repository"
	_userService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImple struct {
	userRepository      _userRepository.UserRepository
	userService         _userService.UserService
	fileStrorageService _fileStorageService.FileStorageService
	conf                *config.Config
}

// Login implements AuthService.
func (a *AuthServiceImple) Login(username, password string) (string, error) {

	user, err := a.userRepository.FindByUsername(username)

	if err != nil {
		log.Println(err)
		return "", err
	}

	if !checkPasswordHash(password, user.Password) {
		return "", &_AuthException.WrongPassword{}
	}

	token, err := a.signToken(user.ID)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil

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

	password, err = hashPassword(password)
	if err != nil {
		return nil, err
	}

	userCreateReq := &_userModel.UserCreateReq{
		Username: username,
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
func (a *AuthServiceImple) VerifyToken(token string) (*_userModel.User, error) {
	hmacSecret := []byte(a.conf.JWT.SecretKey)
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		log.Println(err)
		return nil, &_AuthException.VerifyToken{}
	}
	if !tkn.Valid {
		return nil, &_AuthException.VerifyToken{}
	}
	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &_AuthException.VerifyToken{}
	}
	id, ok := claims["id"].(string)
	if !ok {
		return nil, &_AuthException.VerifyToken{}
	}

	user, err := a.userRepository.FindByID(id)
	if err != nil {
		log.Println(err)
		return nil, &_AuthException.VerifyToken{}
	}

	return user.ToUserModel(), nil
}

func (a *AuthServiceImple) signToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().AddDate(0, 3, 0).Unix(),
	})
	hmacSecret := []byte(a.conf.JWT.SecretKey)
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Println(err)
		return "", &_AuthException.SignToken{}
	}
	return tokenString, nil
}

func getDefaultAvatar() (io.Reader, error) {

	imageURL := "https://source.boringavatars.com/beam/100/"

	// Make HTTP GET request
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, err
	}

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return nil, &_AuthException.DownloadAvatar{}
	}

	imageFileReader := resp.Body

	return imageFileReader, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewAuthServiceImple(userRepository _userRepository.UserRepository, userService _userService.UserService, fileStrorageService _fileStorageService.FileStorageService, conf *config.Config) AuthService {
	return &AuthServiceImple{
		userRepository:      userRepository,
		userService:         userService,
		fileStrorageService: fileStrorageService,
		conf:                conf,
	}
}
