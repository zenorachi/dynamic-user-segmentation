package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
)

type SegmentsService struct {
	segmentsRepo repository.Segments
	usersRepo    repository.Users
}

func NewSegments(segmentsRepo repository.Segments, usersRepo repository.Users) *SegmentsService {
	return &SegmentsService{
		segmentsRepo: segmentsRepo,
		usersRepo:    usersRepo,
	}
}

func (s *SegmentsService) Create(ctx context.Context, segment entity.Segment) (int, error) {
	isExists, err := s.isSegmentExists(ctx, segment.Name)
	if err != nil {
		return 0, err
	}

	if isExists {
		return 0, entity.ErrSegmentAlreadyExists
	}

	if segment.AssignPercent >= 0 {
		segment.AssignPercent /= 100
	} else {
		return 0, entity.ErrInvalidAssignPercent
	}

	return s.segmentsRepo.Create(ctx, segment)
}

func (s *SegmentsService) GetByID(ctx context.Context, id int) (entity.Segment, error) {
	segment, err := s.segmentsRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Segment{}, entity.ErrSegmentDoesNotExist
		}
		return entity.Segment{}, err
	}

	return segment, nil
}

func (s *SegmentsService) GetActiveUsersBySegmentID(ctx context.Context, id int) ([]entity.User, error) {
	_, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.segmentsRepo.GetActiveUsersBySegmentID(ctx, id)
}

func (s *SegmentsService) GetAll(ctx context.Context) ([]entity.Segment, error) {
	segments, err := s.segmentsRepo.GetAll(ctx)
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

	return s.segmentsRepo.DeleteByName(ctx, name)
}

func (s *SegmentsService) DeleteByID(ctx context.Context, id int) error {
	_, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.segmentsRepo.DeleteByID(ctx, id)
}

func (s *SegmentsService) isSegmentExists(ctx context.Context, name string) (bool, error) {
	_, err := s.segmentsRepo.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s *SegmentsService) isUserExists(ctx context.Context, userId int) bool {
	_, err := s.usersRepo.GetByID(ctx, userId)
	return !errors.Is(err, sql.ErrNoRows)
}
