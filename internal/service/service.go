package service

import (
	"context"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/auth"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/hash"
	"time"
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
	}

	Segments interface {
		Create(ctx context.Context, segment entity.Segment) (int, error)
		GetByID(ctx context.Context, id int) (entity.Segment, error)
		GetActiveByUserID(ctx context.Context, userId int) ([]entity.Segment, error)
		GetAll(ctx context.Context) ([]entity.Segment, error)
		DeleteByName(ctx context.Context, name string) error
		DeleteByID(ctx context.Context, id int) error
	}

	Operations interface {
		CreateBySegmentID(ctx context.Context, relations []entity.Relation) ([]int, error)
		CreateBySegmentName(ctx context.Context, userId int, segmentsNames []string) ([]int, error)
		DeleteBySegmentID(ctx context.Context, relations []entity.Relation) ([]int, error)
		DeleteBySegmentName(ctx context.Context, userId int, segmentsNames []string) ([]int, error)
		DeleteAfterTTLBySegmentID(ctx context.Context, relations []entity.Relation, ttl time.Duration)
		DeleteAfterTTLBySegmentName(ctx context.Context, userId int, segmentsNames []string, ttl time.Duration)
	}

	//RelationsTTL interface {
	//	GetTTL(ttl string) (time.Duration, error)
	//	Create(ctx context.Context, relationsTTLs []entity.RelationTTL) error
	//	DeleteAfterTTLBySegmentID(ctx context.Context) error
	//	ScheduleCleanup(ctx context.Context)
	//}
)

type Services struct {
	Users
	Segments
	Operations
	//RelationsTTL
}

type Deps struct {
	Repos           *repository.Repositories
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func New(deps Deps) *Services {
	return &Services{
		Users:      NewUsers(deps.Repos.Users, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL),
		Segments:   NewSegments(deps.Repos.Segments, deps.Repos.Users),
		Operations: NewOperations(deps.Repos.Users, deps.Repos.Segments, deps.Repos.Operations),
		//RelationsTTL: NewRelationsTTL(deps.Repos.RelationsTTL),
	}
}
