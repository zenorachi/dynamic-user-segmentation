package service

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
	"github.com/zenorachi/dynamic-user-segmentation/internal/service/storage"
	"time"
)

type ReportsService struct {
	repo          repository.Operations
	gDriveStorage storage.Provider
}

func NewReports(repo repository.Operations, storage storage.Provider) *ReportsService {
	return &ReportsService{
		repo:          repo,
		gDriveStorage: storage,
	}
}

func (r *ReportsService) CreateReportFile(ctx context.Context, year, month int, userIDs ...int) ([]byte, error) {
	if year < 0 || month < 0 || year > time.Now().Year() || month > 12 {
		return nil, entity.ErrInvalidHistoryPeriod
	}

	operations, err := r.repo.GetOperations(ctx, year, month, userIDs...)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	csvWriter := csv.NewWriter(&buffer)

	header := []string{"user-id", "segment-name", "type", "date"}
	err = csvWriter.Write(header)
	if err != nil {
		return nil, err
	}

	for _, operation := range operations {
		row := []string{
			fmt.Sprintf("%d", operation.UserID), operation.SegmentName, operation.Type,
			operation.Date.Format("2006-01-02 15:04:05"),
		}

		err := csvWriter.Write(row)
		if err != nil {
			return nil, err
		}
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (r *ReportsService) CreateReportLink(ctx context.Context, year, month int, userIDs ...int) (string, error) {
	if !r.gDriveStorage.IsAvailable() {
		return "", entity.ErrGDriveIsNotAvailable
	}

	if year < 0 || month < 0 || year > time.Now().Year() || month > 12 {
		return "", entity.ErrInvalidHistoryPeriod
	}

	data, err := r.CreateReportFile(ctx, year, month, userIDs...)
	if err != nil {
		return "", err
	}

	url, err := r.gDriveStorage.UploadFile(ctx, storage.FileInput{
		Data:        data,
		Name:        fmt.Sprintf("report_%d_%d.csv", year, month),
		ContentType: "text/csv",
	})
	if err != nil {
		return "", err
	}

	return url, nil
}
