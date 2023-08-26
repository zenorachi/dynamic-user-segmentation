package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zenorachi/dynamic-user-segmentation/internal/config"
	"io"
)

type UploadInput struct {
	File        io.Reader
	Name        string
	Size        int64
	ContentType string
}

type Provider interface {
	Upload(ctx context.Context, input UploadInput) (string, error)
	IsAvailable() bool
}

type FileStorage struct {
	client      *minio.Client
	bucket      string
	endpoint    string
	isAvailable bool
}

func NewProvider(cfg *config.MinioConfig) *FileStorage {
	isAvailable := true

	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.User, cfg.Password, ""),
		Secure: false,
	})
	if err != nil {
		isAvailable = false
	}

	return &FileStorage{
		client:      minioClient,
		bucket:      cfg.Bucket,
		endpoint:    cfg.Endpoint,
		isAvailable: isAvailable,
	}
}

func (fs *FileStorage) IsAvailable() bool {
	return fs.isAvailable
}

func (fs *FileStorage) Upload(ctx context.Context, input UploadInput) (string, error) {
	opts := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	if _, err := fs.client.PutObject(ctx, fs.bucket, input.Name, input.File, input.Size, opts); err != nil {
		return "", err
	}

	return fs.generateFileURL(input.Name), nil
}

func (fs *FileStorage) generateFileURL(filename string) string {
	return fmt.Sprintf("http://%s/%s/%s", fs.endpoint, fs.bucket, filename)
}
