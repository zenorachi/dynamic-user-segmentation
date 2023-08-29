package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/zenorachi/dynamic-user-segmentation/internal/config"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type FileInput struct {
	Data        []byte
	Name        string
	ContentType string
}

type Provider interface {
	UploadFile(ctx context.Context, input FileInput) (string, error)
	IsAvailable() bool
}

type GDriveStorage struct {
	gDrive      *drive.Service
	isAvailable bool
}

func NewProvider(cfg *config.GDriveConfig) *GDriveStorage {
	if cfg.Credentials == "" {
		return &GDriveStorage{isAvailable: false}
	}

	gDrive, err := drive.NewService(context.Background(), option.WithCredentialsFile(cfg.Credentials))
	if err != nil {
		logger.Error("google drive", err)
		return &GDriveStorage{isAvailable: false}
	}

	return &GDriveStorage{
		gDrive:      gDrive,
		isAvailable: true,
	}
}

func (g *GDriveStorage) IsAvailable() bool {
	return g.isAvailable
}

func (g *GDriveStorage) UploadFile(ctx context.Context, input FileInput) (string, error) {
	fileId, err := g.getFileIdByName(ctx, input.Name)

	if err != nil {
		if !errors.Is(err, entity.ErrFileNotFound) {
			return "", err
		}

		id, err := g.createFile(ctx, input)
		if err != nil {
			return "", err
		}

		return g.generateFileUrl(id), nil
	}

	err = g.updateFile(ctx, fileId, input.Data)
	if err != nil {
		return "", err
	}

	return g.generateFileUrl(fileId), nil
}

func (g *GDriveStorage) createFile(ctx context.Context, input FileInput) (string, error) {
	file := &drive.File{
		Name:     input.Name,
		MimeType: input.ContentType,
	}

	permissions := &drive.Permission{
		Type: "anyone",
		Role: "reader",
	}

	_, err := g.gDrive.Files.Create(file).Context(ctx).Media(bytes.NewReader(input.Data)).Do()
	if err != nil {
		return "", err
	}

	fileID, err := g.getFileIdByName(ctx, input.Name)
	if err != nil {
		return "", err
	}

	_, err = g.gDrive.Permissions.Create(fileID, permissions).Context(ctx).Do()
	if err != nil {
		return "", err
	}

	return fileID, nil
}

func (g *GDriveStorage) updateFile(ctx context.Context, id string, data []byte) error {
	_, err := g.gDrive.Files.Update(id, &drive.File{}).Context(ctx).Media(bytes.NewReader(data)).Do()

	return err
}

func (g *GDriveStorage) getFileIdByName(ctx context.Context, name string) (string, error) {
	files, err := g.getAllFiles(ctx)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.Name == name {
			return file.Id, nil
		}
	}

	return "", entity.ErrFileNotFound
}

func (g *GDriveStorage) getAllFiles(ctx context.Context) ([]*drive.File, error) {
	r, err := g.gDrive.Files.List().Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return r.Files, nil
}

func (g *GDriveStorage) generateFileUrl(id string) string {
	return fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", id)
}
