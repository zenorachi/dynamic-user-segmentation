package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
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

func (o *OperationsService) CreateRelationsBySegmentIDs(ctx context.Context, userId int, segmentIDs []int) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.CreateRelationsBySegmentIDs(ctx, userId, segmentIDs)
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

func (o *OperationsService) CreateRelationsBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.CreateRelationsBySegmentNames(ctx, userId, segmentsNames)
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

func (o *OperationsService) DeleteRelationsBySegmentIDs(ctx context.Context, userId int, segmentIDs []int) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.DeleteRelationsBySegmentIDs(ctx, userId, segmentIDs)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrSegmentDoesNotExist
		}
		return nil, err
	}

	return operations, nil
}

func (o *OperationsService) DeleteRelationsBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error) {
	if !o.isUserExists(ctx, userId) {
		return nil, entity.ErrUserDoesNotExist
	}

	operations, err := o.operationsRepo.DeleteRelationsBySegmentNames(ctx, userId, segmentsNames)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrSegmentDoesNotExist
		}
		return nil, err
	}

	return operations, nil
}

func (o *OperationsService) DeleteRelationsAfterTTLBySegmentIDs(ctx context.Context, userId int, segmentsIDs []int, ttl time.Duration) {
	select {
	case <-time.After(ttl):
		_, err := o.DeleteRelationsBySegmentIDs(ctx, userId, segmentsIDs)
		if err != nil {
			logger.Error("scheduler", err)
		} else {
			logger.Info("scheduler", "success")
		}
	case <-ctx.Done():
		return
	}
}

func (o *OperationsService) DeleteRelationsAfterTTLBySegmentNames(ctx context.Context, userId int, segmentsNames []string, ttl time.Duration) {
	select {
	case <-time.After(ttl):
		_, err := o.DeleteRelationsBySegmentNames(ctx, userId, segmentsNames)
		if err != nil {
			logger.Error("scheduler", err)
		} else {
			logger.Info("scheduler", "success")
		}
	case <-ctx.Done():
		return
	}
}

func (o *OperationsService) GetOperationsHistory(ctx context.Context, year, month int, userIDs ...int) ([]entity.Operation, error) {
	if year < 0 || month < 0 || year > time.Now().Year() || month > 12 {
		return nil, entity.ErrInvalidHistoryPeriod
	}
	return o.operationsRepo.GetOperations(ctx, year, month, userIDs...)
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
