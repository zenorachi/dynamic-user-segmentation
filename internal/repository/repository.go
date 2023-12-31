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
		GetActiveSegmentsByUserID(ctx context.Context, id int) ([]entity.Segment, error)
		SetSession(ctx context.Context, userId int, session entity.Session) error
	}

	Segments interface {
		Create(ctx context.Context, segment entity.Segment) (int, error)
		GetByName(ctx context.Context, name string) (entity.Segment, error)
		GetByID(ctx context.Context, id int) (entity.Segment, error)
		GetActiveUsersBySegmentID(ctx context.Context, id int) ([]entity.User, error)
		GetAll(ctx context.Context) ([]entity.Segment, error)
		DeleteByName(ctx context.Context, name string) error
		DeleteByID(ctx context.Context, id int) error
	}

	Operations interface {
		CreateRelationsBySegmentIDs(ctx context.Context, userId int, segmentIDs []int) ([]int, error)
		CreateRelationsBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error)
		DeleteRelationsBySegmentIDs(ctx context.Context, userId int, segmentsIDs []int) ([]int, error)
		DeleteRelationsBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error)
		GetOperations(ctx context.Context, year, month int, userIDs ...int) ([]entity.Operation, error)
	}
)

type Repositories struct {
	Users
	Segments
	Operations
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		Users:      NewUsers(db),
		Segments:   NewSegments(db),
		Operations: NewOperations(db),
	}
}
