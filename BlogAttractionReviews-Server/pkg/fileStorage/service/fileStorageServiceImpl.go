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
func (f *FileStorageServiceImpl) GetFile(filename string) (io.Reader, error) {
	file, err := f.fileStorageRepository.GetFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// UploadFile implements FileStorageService.
func (f *FileStorageServiceImpl) UploadFile(file io.Reader, filename string) (string, error) {
	filename, err := f.fileStorageRepository.UploadFile(file, filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func NewFileStorageServiceImpl(fileStorageRepository _fileStorageRepository.FileStorageRepository) FileStorageService {
	return &FileStorageServiceImpl{fileStorageRepository: fileStorageRepository}
}
