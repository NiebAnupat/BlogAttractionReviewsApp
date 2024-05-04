package repository

import (
	"io"
	"sync"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileStorageRepositoryImpl struct {
	conf *config.Config
}

var (
	s3Client *s3.S3
	once     sync.Once
)

func (f *FileStorageRepositoryImpl) getS3Client() *s3.S3 {

	once.Do(func() {
		s3Client = s3.New(session.Must(session.NewSession(&aws.Config{
			Region:      aws.String(f.conf.AWS.Region),
			Credentials: credentials.NewStaticCredentials(f.conf.AWS.S3.AccessKeyID, f.conf.AWS.S3.SecretAccessKey, ""),
		})))
	})
	return s3Client
}

// DeleteFile implements FileStorageRepository.
func (f *FileStorageRepositoryImpl) DeleteFile(filename string) error {
	s3Client := f.getS3Client()

	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(f.conf.AWS.S3.Bucket),
		Key:    aws.String(filename),
	})

	if err != nil {
		return err
	}

	return nil
}

// GetFile implements FileStorageRepository.
func (f *FileStorageRepositoryImpl) GetFile(filename string) ([]byte, error) {
	panic("unimplemented")
}

// UploadFile implements FileStorageRepository.
func (f *FileStorageRepositoryImpl) UploadFile(file io.Reader, filename string) (string, error) {

	s3Client := f.getS3Client()

	uploader := s3manager.NewUploaderWithClient(s3Client)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(f.conf.AWS.S3.Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return filename, nil
}

func NewFileStorageRepositoryImpl(conf *config.Config) FileStorageRepository {
	return &FileStorageRepositoryImpl{conf: conf}
}
