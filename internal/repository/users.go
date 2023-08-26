package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) Create(ctx context.Context, user entity.User) (int, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return 0, err
	}

	var (
		id    int
		query = fmt.Sprintf("INSERT INTO %s (login, email, password) VALUES ($1, $2, $3) RETURNING id",
			collectionUsers)
	)

	err = tx.QueryRowContext(ctx, query, user.Login, user.Email, user.Password).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (u *UsersRepository) GetByID(ctx context.Context, id int) (entity.User, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.User{}, err
	}

	var (
		user  entity.User
		query = fmt.Sprintf("SELECT id, login, email, password, registered_at FROM %s WHERE id = $1",
			collectionUsers)
	)

	err = tx.QueryRowContext(ctx, query, id).
		Scan(&user.ID, &user.Login, &user.Email, &user.Password, &user.RegisteredAt)
	if err != nil {
		_ = tx.Rollback()
		return entity.User{}, err
	}

	return user, tx.Commit()
}

func (u *UsersRepository) GetByLogin(ctx context.Context, login string) (entity.User, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.User{}, err
	}

	var (
		user  entity.User
		query = fmt.Sprintf("SELECT id, login, email, password, registered_at FROM %s WHERE login = $1",
			collectionUsers)
	)

	err = tx.QueryRowContext(ctx, query, login).
		Scan(&user.ID, &user.Login, &user.Email, &user.Password, &user.RegisteredAt)
	if err != nil {
		_ = tx.Rollback()
		return entity.User{}, err
	}

	return user, tx.Commit()
}

func (u *UsersRepository) GetByCredentials(ctx context.Context, login, password string) (entity.User, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.User{}, err
	}

	var (
		user  entity.User
		query = fmt.Sprintf("SELECT id, login, email, password, registered_at FROM %s WHERE login = $1 AND password = $2",
			collectionUsers)
	)

	err = tx.QueryRowContext(ctx, query, login, password).
		Scan(&user.ID, &user.Login, &user.Email, &user.Password, &user.RegisteredAt)
	if err != nil {
		_ = tx.Rollback()
		return entity.User{}, err
	}

	return user, tx.Commit()
}

func (u *UsersRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (entity.User, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.User{}, err
	}

	var (
		user  entity.User
		query = fmt.Sprintf("SELECT id, login, email, password, registered_at FROM %s WHERE (session).\"refresh_token\" = $1",
			collectionUsers)
	)

	err = tx.QueryRowContext(ctx, query, refreshToken).
		Scan(&user.ID, &user.Login, &user.Email, &user.Password, &user.RegisteredAt)
	if err != nil {
		_ = tx.Rollback()
		return entity.User{}, err
	}

	return user, tx.Commit()
}

func (u *UsersRepository) SetSession(ctx context.Context, userId int, session entity.Session) error {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET session = ROW($1, $2) WHERE id = $3",
		collectionUsers)

	_, err = tx.ExecContext(ctx, query, session.RefreshToken, session.ExpiresAt, userId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
