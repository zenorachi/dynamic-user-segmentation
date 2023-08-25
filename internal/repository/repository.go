package repository

import (
	"context"
	"database/sql"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

type (
	Users interface {
		Create(ctx context.Context, user entity.User) (int, error)
		GetByID(ctx context.Context, id int) (entity.User, error)
		GetByLogin(ctx context.Context, login string) (entity.User, error)
		GetByCredentials(ctx context.Context, login, password string) (entity.User, error)
		GetByRefreshToken(ctx context.Context, refreshToken string) (entity.User, error)
		SetSession(ctx context.Context, userId int, session entity.Session) error
	}

	Segments interface {
		Create(ctx context.Context, segment entity.Segment) (int, error)
		GetByName(ctx context.Context, name string) (entity.Segment, error)
		GetByID(ctx context.Context, id int) (entity.Segment, error)
		GetAll(ctx context.Context) ([]entity.Segment, error)
		DeleteByName(ctx context.Context, name string) error
		DeleteByID(ctx context.Context, id int) error
	}

	//Relations interface {
	//	CreateBySegmentID(ctx context.Context, relation entity.Relation) (int, error)
	//	CreateBySegmentName(ctx context.Context, userId int, segmentName string) (int, error)
	//	GetByUserID(ctx context.Context, userId int) (entity.Relation, error)
	//	GetBySegmentID(ctx context.Context, segmentId int) (entity.Relation, error)
	//	DeleteBySegmentID(ctx context.Context, relation entity.Relation) (int, error)
	//	DeleteBySegmentName(ctx context.Context, userId int, segmentName string) (int, error)
	//}

	Operations interface {
		CreateBySegmentID(ctx context.Context, relation entity.Relation) (int, error)
		CreateBySegmentName(ctx context.Context, userId int, segmentName string) (int, error)
		GetByBothIDs(ctx context.Context, userId, segmentId int) (entity.Relation, error)
		//TODO INVALID RETURN PARAMETER
		GetByUserID(ctx context.Context, userId int) (entity.Relation, error)
		GetBySegmentID(ctx context.Context, segmentId int) (entity.Relation, error)
		DeleteBySegmentID(ctx context.Context, relation entity.Relation) (int, error)
		DeleteBySegmentName(ctx context.Context, userId int, segmentName string) (int, error)
		GetOperations(ctx context.Context, userIds ...int) ([]entity.Operation, error)
		//GetByUserID(ctx context.Context, userId int) ([]entity.Operation, error)
		//GetAll(ctx context.Context) ([]entity.Operation, error)
	}
)

type Repositories struct {
	Users
	Segments
	//Relations
	Operations
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		Users:    NewUsers(db),
		Segments: NewSegments(db),
		//Relations:  NewRelations(db),
		Operations: NewOperations(db),
	}
}
