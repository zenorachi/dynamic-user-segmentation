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
		GetAll(ctx context.Context) ([]entity.Segment, error)
		DeleteByName(ctx context.Context, name string) error
		DeleteByID(ctx context.Context, id int) error
	}

	Operations interface {
		CreateBySegmentID(ctx context.Context, relation entity.Relation) (int, error)
		CreateBySegmentName(ctx context.Context, userId int, segmentName string) (int, error)
		DeleteBySegmentID(ctx context.Context, relation entity.Relation) (int, error)
		DeleteBySegmentName(ctx context.Context, userId int, segmentName string) (int, error)
	}
)

type Services struct {
	Users
	Segments
	Operations
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
		Segments:   NewSegments(deps.Repos.Segments),
		Operations: NewOperations(deps.Repos.Users, deps.Repos.Segments, deps.Repos.Operations),
	}
}
