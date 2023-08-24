package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/auth"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/hash"
	"time"
)

type UserService struct {
	repo            repository.Users
	hasher          hash.PasswordHasher
	tokenManager    auth.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUsers(repo repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) *UserService {
	return &UserService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (u *UserService) SignUp(ctx context.Context, login, email, password string) (int, error) {
	if u.isUserExists(ctx, login) {
		return 0, entity.ErrUserAlreadyExists
	}

	hashedPassword, err := u.hasher.Hash(password)
	if err != nil {
		return 0, err
	}

	user := entity.NewUser(login, email, hashedPassword)
	return u.repo.Create(ctx, user)
}

func (u *UserService) SignIn(ctx context.Context, login, password string) (Tokens, error) {
	if !u.isUserExists(ctx, login) {
		return Tokens{}, entity.ErrUserDoesNotExist
	}
	hashedPassword, err := u.hasher.Hash(password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := u.repo.GetByCredentials(ctx, login, hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Tokens{}, entity.ErrIncorrectPassword
		}
		return Tokens{}, err
	}

	return u.createSession(ctx, user.ID)
}

func (u *UserService) RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error) {
	user, err := u.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Tokens{}, entity.ErrSessionDoesNotExist
		}
		return Tokens{}, err
	}
	return u.createSession(ctx, user.ID)
}

func (u *UserService) createSession(ctx context.Context, userId int) (Tokens, error) {
	var (
		tokens Tokens
		err    error
	)

	tokens.AccessToken, err = u.tokenManager.NewJWT(userId, u.accessTokenTTL)
	if err != nil {
		return Tokens{}, err
	}

	tokens.RefreshToken, err = u.tokenManager.NewRefreshToken()
	if err != nil {
		return Tokens{}, err
	}

	session := entity.NewSession(tokens.RefreshToken, time.Now().Add(u.refreshTokenTTL))
	return tokens, u.repo.SetSession(ctx, userId, session)
}

func (u *UserService) isUserExists(ctx context.Context, login string) bool {
	_, err := u.repo.GetByLogin(ctx, login)
	return !errors.Is(err, sql.ErrNoRows)
}
