package storage

import (
	"context"
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
}

func NewUploadInput(file io.Reader, name string, size int64, contentType string) UploadInput {
	return UploadInput{
		File:        file,
		Name:        name,
		Size:        size,
		ContentType: contentType,
	}
}
