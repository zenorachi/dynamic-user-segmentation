package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
)

type SegmentsService struct {
	repo repository.Segments
}

func NewSegments(repo repository.Segments) *SegmentsService {
	return &SegmentsService{repo: repo}
}

func (s *SegmentsService) Create(ctx context.Context, segment entity.Segment) (int, error) {
	isExists, err := s.isSegmentExists(ctx, segment.Name)
	if err != nil {
		return 0, err
	}

	if isExists {
		return 0, entity.ErrSegmentAlreadyExists
	}

	return s.repo.Create(ctx, segment)
}

func (s *SegmentsService) isSegmentExists(ctx context.Context, name string) (bool, error) {
	_, err := s.repo.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
