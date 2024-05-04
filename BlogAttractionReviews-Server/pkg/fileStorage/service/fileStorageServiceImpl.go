package service

import (
	"io"

	_fileStorageRepository "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/fileStorage/repository"
)

type FileStorageServiceImpl struct {
	fileStorageRepository _fileStorageRepository.FileStorageRepository
}

// DeleteFile implements FileStorageService.
func (f *FileStorageServiceImpl) DeleteFile(filename string) error {
	err := f.fileStorageRepository.DeleteFile(filename)
	if err != nil {
		return err
	}
	return nil
}

// GetFile implements FileStorageService.
func (f *FileStorageServiceImpl) GetFile(filename string) ([]byte, error) {
	panic("unimplemented")
}

// UploadFile implements FileStorageService.
func (f *FileStorageServiceImpl) UploadFile(file io.Reader, filename string) (string, error) {
	imageURL, err := f.fileStorageRepository.UploadFile(file, filename)
	if err != nil {
		return "", err
	}
	return imageURL, nil
}

func NewFileStorageServiceImpl(fileStorageRepository _fileStorageRepository.FileStorageRepository) FileStorageService {
	return &FileStorageServiceImpl{fileStorageRepository: fileStorageRepository}
}
