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

func (o *OperationsService) CreateBySegmentID(ctx context.Context, relations []entity.Relation) ([]int, error) {
	if !o.isUserExists(ctx, relations[0].UserID) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.CreateBySegmentID(ctx, relations)
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

func (o *OperationsService) CreateBySegmentName(ctx context.Context, userId int, segmentsNames []string) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.CreateBySegmentName(ctx, userId, segmentsNames)
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

func (o *OperationsService) DeleteBySegmentID(ctx context.Context, relations []entity.Relation) ([]int, error) {
	if !o.isUserExists(ctx, relations[0].UserID) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.DeleteBySegmentID(ctx, relations)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrSegmentDoesNotExist
		}
		return nil, err
	}

	return operations, nil
}

func (o *OperationsService) DeleteBySegmentName(ctx context.Context, userId int, segmentsNames []string) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.DeleteBySegmentName(ctx, userId, segmentsNames)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrSegmentDoesNotExist
		}
		return nil, err
	}

	return operations, nil
}

func (o *OperationsService) DeleteAfterTTLBySegmentID(ctx context.Context, relations []entity.Relation, ttl time.Duration) {
	select {
	case <-time.After(ttl):
		_, err := o.DeleteBySegmentID(ctx, relations)
		if err != nil {
			logger.Error("scheduler", err)
		} else {
			logger.Info("scheduler", "success")
		}
	case <-ctx.Done():
		return
	}
}

func (o *OperationsService) DeleteAfterTTLBySegmentName(ctx context.Context, userId int, segmentsNames []string, ttl time.Duration) {
	select {
	case <-time.After(ttl):
		_, err := o.DeleteBySegmentName(ctx, userId, segmentsNames)
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
