package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
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

func (o *OperationsService) CreateBySegmentID(ctx context.Context, relation entity.Relation) (int, error) {
	_, err := o.usersRepo.GetByID(ctx, relation.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, entity.ErrUserDoesNotExist
		}
	}

	_, err = o.segmentsRepo.GetByID(ctx, relation.SegmentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, entity.ErrSegmentDoesNotExist
		}
	}

	isExists, err := o.isRelationExists(ctx, relation.UserID, relation.SegmentID)
	if err != nil {
		return 0, err
	}
	if isExists {
		return 0, entity.ErrRelationAlreadyExists
	}

	return o.operationsRepo.CreateBySegmentID(ctx, relation)
}

func (o *OperationsService) CreateBySegmentName(ctx context.Context, userId int, segmentName string) (int, error) {
	//_, err := o.usersRepo.GetByID(ctx, userId)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return 0, entity.ErrUserDoesNotExist
	//	}
	//}
	//
	//_, err = o.segmentsRepo.GetByName(ctx, segmentName)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return 0, entity.ErrSegmentDoesNotExist
	//	}
	//}
	//
	//isExists, err := o.isRelationExists(ctx, userId)
	//if err != nil {
	//	return 0, err
	//}
	//if isExists {
	//	return 0, entity.ErrRelationAlreadyExists
	//}

	return o.operationsRepo.CreateBySegmentName(ctx, userId, segmentName)
}

func (o *OperationsService) DeleteBySegmentID(ctx context.Context, relation entity.Relation) (int, error) {
	isExists, err := o.isRelationExists(ctx, relation.UserID, relation.SegmentID)
	if err != nil {
		return 0, err
	}
	if !isExists {
		return 0, entity.ErrRelationDoesNotExist
	}

	return o.operationsRepo.DeleteBySegmentID(ctx, relation)
}

func (o *OperationsService) DeleteBySegmentName(ctx context.Context, userId int, segmentName string) (int, error) {
	//isExists, err := o.isRelationExists(ctx, userId)
	//if err != nil {
	//	return 0, err
	//}
	//if !isExists {
	//	return 0, entity.ErrRelationDoesNotExist
	//}

	return o.operationsRepo.DeleteBySegmentName(ctx, userId, segmentName)
}

func (o *OperationsService) isRelationExists(ctx context.Context, userId, segmentId int) (bool, error) {
	_, err := o.operationsRepo.GetByBothIDs(ctx, userId, segmentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
