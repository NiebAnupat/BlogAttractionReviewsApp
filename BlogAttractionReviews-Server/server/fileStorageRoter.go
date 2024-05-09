package server

import (
	_controller "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/controller"
	_repository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/repository"
	_service "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/service"
)

func (f *fiberServer) initFileStorageRouter() {
	fileStorageRepository := _repository.NewFileStorageRepositoryImpl(f.conf)
	fileStorageService := _service.NewFileStorageServiceImpl(fileStorageRepository)
	fileStorageController := _controller.NewFileStorageController(fileStorageService)

	router := f.app.Group("/v1/fileStorage")
	router.Get("/:filename", fileStorageController.GetFile)
	// router.Delete("/:filename", fileStorageController.DeleteFile)
	// router.Post("/", fileStorageController.UploadFile)
}
