package service

import (
	"context"
	"time"

	"github.com/zenorachi/dynamic-user-segmentation/internal/service/storage"

	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/auth"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/hash"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type (
	Users interface {
		SignUp(ctx context.Context, login, email, password string) (int, error)
		SignIn(ctx context.Context, login, password string) (Tokens, error)
		RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
		GetActiveSegmentsByUserID(ctx context.Context, id int) ([]entity.Segment, error)
	}

	Segments interface {
		Create(ctx context.Context, segment entity.Segment) (int, error)
		GetByID(ctx context.Context, id int) (entity.Segment, error)
		GetActiveUsersBySegmentID(ctx context.Context, id int) ([]entity.User, error)
		GetAll(ctx context.Context) ([]entity.Segment, error)
		DeleteByName(ctx context.Context, name string) error
		DeleteByID(ctx context.Context, id int) error
	}

	Operations interface {
		CreateRelationsBySegmentIDs(ctx context.Context, userId int, segmentIDs []int) ([]int, error)
		CreateRelationsBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error)
		DeleteRelationsBySegmentIDs(ctx context.Context, userId int, segmentIDs []int) ([]int, error)
		DeleteRelationsBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error)
		DeleteRelationsAfterTTLBySegmentIDs(ctx context.Context, userId int, segmentIDs []int, ttl time.Duration)
		DeleteRelationsAfterTTLBySegmentNames(ctx context.Context, userId int, segmentsNames []string, ttl time.Duration)
		GetOperationsHistory(ctx context.Context, year, month int, userIDs ...int) ([]entity.Operation, error)
	}

	Reports interface {
		CreateReportFile(ctx context.Context, year, month int, userIDs ...int) ([]byte, error)
		CreateReportLink(ctx context.Context, year, month int, userIDs ...int) (string, error)
	}
)

type Services struct {
	Users
	Segments
	Operations
	Reports
}

type Deps struct {
	Repos           *repository.Repositories
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	Storage         *storage.GDriveStorage
}

func New(deps Deps) *Services {
	return &Services{
		Users:      NewUsers(deps.Repos.Users, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL),
		Segments:   NewSegments(deps.Repos.Segments, deps.Repos.Users),
		Operations: NewOperations(deps.Repos.Users, deps.Repos.Segments, deps.Repos.Operations),
		Reports:    NewReports(deps.Repos.Operations, deps.Storage),
	}
}
