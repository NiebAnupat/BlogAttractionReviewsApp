package repository

type FileStorageRepository interface {
	UploadFile(file []byte, filename string) (string, error)
	GetFile(filename string) ([]byte, error)
	DeleteFile(filename string) error
	// UploadFileWithBucket(file []byte, filename, bucket string) (string, error)
}
