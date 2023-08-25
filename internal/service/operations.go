package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
	"time"
)

type OperationsService struct {
	usersRepo      repository.Users
	segmentsRepo   repository.Segments
	operationsRepo repository.Operations
}

func NewOperations(usersRepo repository.Users, segmentsRepo repository.Segments,
	operationsRepo repository.Operations) *OperationsService {
	return &OperationsService{
		usersRepo:      usersRepo,
		segmentsRepo:   segmentsRepo,
		operationsRepo: operationsRepo,
	}
}

func (o *OperationsService) CreateBySegmentIDs(ctx context.Context, userId int, segmentIDs []int) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.CreateBySegmentIDs(ctx, userId, segmentIDs)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrSegmentDoesNotExist
		}
		if o.isRelationExistsError(err) {
			return nil, entity.ErrRelationAlreadyExists
		}
		return nil, err
	}

	return operations, nil
}

func (o *OperationsService) CreateBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.CreateBySegmentNames(ctx, userId, segmentsNames)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrSegmentDoesNotExist
		}
		if o.isRelationExistsError(err) {
			return nil, entity.ErrRelationAlreadyExists
		}
		return nil, err
	}

	return operations, nil
}

func (o *OperationsService) DeleteBySegmentIDs(ctx context.Context, userId int, segmentIDs []int) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.DeleteBySegmentIDs(ctx, userId, segmentIDs)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrSegmentDoesNotExist
		}
		return nil, err
	}

	return operations, nil
}

func (o *OperationsService) DeleteBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.DeleteBySegmentNames(ctx, userId, segmentsNames)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrSegmentDoesNotExist
		}
		return nil, err
	}

	return operations, nil
}

func (o *OperationsService) DeleteAfterTTLBySegmentIDs(ctx context.Context, userId int, segmentsIDs []int, ttl time.Duration) {
	select {
	case <-time.After(ttl):
		_, err := o.DeleteBySegmentIDs(ctx, userId, segmentsIDs)
		if err != nil {
			logger.Error("scheduler", err)
		} else {
			logger.Info("scheduler", "success")
		}
	case <-ctx.Done():
		return
	}
}

func (o *OperationsService) DeleteAfterTTLBySegmentNames(ctx context.Context, userId int, segmentsNames []string, ttl time.Duration) {
	select {
	case <-time.After(ttl):
		_, err := o.DeleteBySegmentNames(ctx, userId, segmentsNames)
		if err != nil {
			logger.Error("scheduler", err)
		} else {
			logger.Info("scheduler", "success")
		}
	case <-ctx.Done():
		return
	}
}

func (o *OperationsService) isUserExists(ctx context.Context, userId int) bool {
	_, err := o.usersRepo.GetByID(ctx, userId)
	return !errors.Is(err, sql.ErrNoRows)
}

func (o *OperationsService) isRelationExistsError(err error) bool {
	var pqErr *pq.Error
	isPqError := errors.As(err, &pqErr)

	if isPqError && pqErr.Code == "23505" {
		return true
	}

	return false
}
