package repository

import "io"

type FileStorageRepository interface {
	UploadFile(file io.Reader, filename string) (string, error)
	GetFile(filename string) (io.Reader, error)
	DeleteFile(filename string) error
	// UploadFileWithBucket(file []byte, filename, bucket string) (string, error)
}
