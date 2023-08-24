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

func (s *SegmentsService) GetByID(ctx context.Context, id int) (entity.Segment, error) {
	segment, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Segment{}, entity.ErrSegmentDoesNotExist
		}
		return entity.Segment{}, err
	}

	return segment, nil
}

func (s *SegmentsService) GetAll(ctx context.Context) ([]entity.Segment, error) {
	segments, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(segments) == 0 {
		return nil, entity.ErrSegmentDoesNotExist
	}

	return segments, nil
}

func (s *SegmentsService) DeleteByName(ctx context.Context, name string) error {
	isExists, err := s.isSegmentExists(ctx, name)
	if err != nil {
		return err
	}

	if !isExists {
		return entity.ErrSegmentDoesNotExist
	}

	return s.repo.DeleteByName(ctx, name)
}

func (s *SegmentsService) DeleteByID(ctx context.Context, id int) error {
	_, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.DeleteByID(ctx, id)
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
