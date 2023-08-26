package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zenorachi/dynamic-user-segmentation/internal/config"
	"os"
)

func NewProvider(cfg *config.Config) (*minio.Client, error) {
	minioRootUser := os.Getenv("MINIO_ROOT_USER")
	minioRootPassword := os.Getenv("MINIO_ROOT_PASSWORD")

	minioClient, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioRootUser, minioRootPassword, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
