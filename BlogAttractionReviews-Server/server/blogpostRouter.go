package server

import (
	_blogPostController "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/controller"
	_blogPostRepository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/repository"
	_blogPostService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/service"

	_authService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/auth/service"
	_fileStorageRepository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/repository"
	_fileStorageService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/service"
	_userRepository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/repository"
	_userService "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/service"
)

func (f *fiberServer) initBlogPostRouter() {

	authRepository := _userRepository.NewUserRepositoryImpl(f.db)
	fileStoreageRepository := _fileStorageRepository.NewFileStorageRepositoryImpl(f.conf)
	fileStorageService := _fileStorageService.NewFileStorageServiceImpl(fileStoreageRepository)
	userRepository := _userRepository.NewUserRepositoryImpl(f.db)
	userService := _userService.NewUserServiceImpl(authRepository, fileStorageService)
	authService := _authService.NewAuthServiceImple(userRepository, userService, fileStorageService, f.conf)

	blogPostRepository := _blogPostRepository.NewBlogPostRepositoryImpl(f.db)
	blogPostService := _blogPostService.NewBlogPostServiceImpl(blogPostRepository)
	blogPostController := _blogPostController.NewBlogPostController(blogPostService, authService, fileStorageService)

	router := f.app.Group("/v1/blog")
	router.Post("/", blogPostController.PostNewBlog)
	router.Post("/content", blogPostController.AddContentToBlogPost)

}
