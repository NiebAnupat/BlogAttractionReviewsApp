package server

import (
	_authController "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/controller"
	_authService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/service"
	_fileStorageRepository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/repository"
	_fileStorageService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/service"
	_userRepository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/repository"
	_userService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/service"
)

func (f *fiberServer) initAuthRouter() {
	authRepository := _userRepository.NewUserRepositoryImpl(f.db)
	fileStoreageRepository := _fileStorageRepository.NewFileStorageRepositoryImpl(f.conf)
	fileStorageService := _fileStorageService.NewFileStorageServiceImpl(fileStoreageRepository)
	userService := _userService.NewUserServiceImpl(authRepository, fileStorageService)
	authService := _authService.NewAuthServiceImple(userService, fileStorageService)
	authController := _authController.NewAuthControllerImpl(authService)

	router := f.app.Group("/v1/auth")

	router.Post("/register", authController.Register)
	// router.Post("/login", authController.Login)
	// router.Post("/logout", authController.Logout)
	// router.Post("/verify", authController.VerifyToken)
	// router.Post("/refresh", authController.RefreshToken)
}
