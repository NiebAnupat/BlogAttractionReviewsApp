package service

import "io"

type FileStorageService interface {
	UploadFile(file io.Reader, filename string) (string, error)
	GetFile(filename string) ([]byte, error)
	DeleteFile(filename string) error
}
