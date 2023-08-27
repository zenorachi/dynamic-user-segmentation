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

func (u *UsersRepository) GetActiveSegmentsByUserID(ctx context.Context, id int) ([]entity.Segment, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return nil, err
	}

	var (
		segments []entity.Segment
		query    = fmt.Sprintf(
			"SELECT name FROM %s JOIN %s ON %s.segment_id = id WHERE %s.user_id = $1",
			collectionSegments, collectionRelations, collectionRelations, collectionRelations)
	)

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var segment entity.Segment
		err = rows.Scan(&segment.Name)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		segments = append(segments, segment)
	}

	if err = rows.Err(); err != nil {
		_ = tx.Commit()
		return nil, err
	}

	return segments, tx.Commit()
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

//func (u *UsersRepository) AutoAssignUsers(ctx context.Context) {
//	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  false,
//	})
//	if err != nil {
//		logger.Error("auto assign", err)
//		return
//	}
//
//	usersCount, err := u.getTotalUsers(ctx)
//	if err != nil {
//		_ = tx.Rollback()
//		logger.Error("auto assign", err)
//		return
//	}
//
//	segments, err := u.getSegments(ctx)
//	if err != nil {
//		_ = tx.Rollback()
//		logger.Error("auto assign", err)
//		return
//	}
//
//	var (
//		queryInsertRelation = fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2)",
//			collectionRelations)
//		queryInsertOperations = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3)",
//			collectionOperations)
//	)
//
//	for _, segment := range segments {
//		limit := float64(usersCount) * segment.AssignPercent
//		userIDs, err := u.getRandomUserIDs(ctx, int(limit))
//		if err != nil {
//			_ = tx.Rollback()
//			logger.Error("auto assign", err)
//			return
//		}
//
//		for _, id := range userIDs {
//			_, err = tx.ExecContext(ctx, queryInsertRelation, id, segment.ID)
//			if err != nil && !u.isRelationExistsError(err) {
//				_ = tx.Rollback()
//				logger.Error("auto assign", err)
//				return
//			}
//
//			_, err = tx.ExecContext(ctx, queryInsertOperations, id, segment.Name, "added")
//			if err != nil {
//				_ = tx.Rollback()
//				logger.Error("auto assign", err)
//				return
//			}
//		}
//	}
//
//	_ = tx.Commit()
//}
//
//func (u *UsersRepository) getTotalUsers(ctx context.Context) (int, error) {
//	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  true,
//	})
//
//	var (
//		userCount       int
//		queryCountUsers = fmt.Sprintf("SELECT COUNT(*) FROM %s", collectionUsers)
//	)
//
//	err = tx.QueryRow(queryCountUsers).Scan(&userCount)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	return userCount, tx.Commit()
//}
//
//func (u *UsersRepository) getSegments(ctx context.Context) ([]entity.Segment, error) {
//	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  true,
//	})
//
//	var (
//		segments         []entity.Segment
//		queryGetSegments = fmt.Sprintf("SELECT * FROM %s WHERE assign_percent > 0.0",
//			collectionSegments)
//	)
//
//	rows, err := tx.Query(queryGetSegments)
//	if err != nil {
//		_ = tx.Rollback()
//		return nil, err
//	}
//
//	for rows.Next() {
//		var segment entity.Segment
//		err = rows.Scan(&segment.ID, &segment.Name, &segment.AssignPercent)
//		if err != nil {
//			_ = tx.Rollback()
//			return nil, err
//		}
//
//		segments = append(segments, segment)
//	}
//
//	return segments, tx.Commit()
//}
//
//func (u *UsersRepository) getRandomUserIDs(ctx context.Context, limit int) ([]int, error) {
//	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  true,
//	})
//
//	var (
//		userIDs          []int
//		queryGetSegments = fmt.Sprintf("SELECT id FROM %s ORDER BY random() LIMIT %d",
//			collectionSegments, limit)
//	)
//
//	rows, err := tx.Query(queryGetSegments)
//	if err != nil {
//		_ = tx.Rollback()
//		return nil, err
//	}
//
//	for rows.Next() {
//		var id int
//		err = rows.Scan(&id)
//		if err != nil {
//			_ = tx.Rollback()
//			return nil, err
//		}
//
//		userIDs = append(userIDs, id)
//	}
//
//	return userIDs, tx.Commit()
//}
//
//func (u *UsersRepository) isRelationExistsError(err error) bool {
//	var pqErr *pq.Error
//	isPqError := errors.As(err, &pqErr)
//
//	if isPqError && pqErr.Code == "23505" {
//		return true
//	}
//
//	return false
//}
