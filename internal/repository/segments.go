package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

type SegmentsRepository struct {
	db *sql.DB
}

func NewSegments(db *sql.DB) *SegmentsRepository {
	return &SegmentsRepository{db: db}
}

func (s *SegmentsRepository) Create(ctx context.Context, segment entity.Segment) (int, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return 0, err
	}

	var (
		id    int
		query = fmt.Sprintf("INSERT INTO %s (name, assign_percent) VALUES ($1, $2) RETURNING id",
			collectionSegments)
	)

	err = tx.QueryRowContext(ctx, query, segment.Name, segment.AssignPercent).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (s *SegmentsRepository) GetByName(ctx context.Context, name string) (entity.Segment, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.Segment{}, err
	}

	var (
		segment entity.Segment
		query   = fmt.Sprintf("SELECT id, name FROM %s WHERE name = $1",
			collectionSegments)
	)

	err = tx.QueryRowContext(ctx, query, name).
		Scan(&segment.ID, &segment.Name)
	if err != nil {
		_ = tx.Rollback()
		return entity.Segment{}, err
	}

	return segment, tx.Commit()
}

func (s *SegmentsRepository) GetByID(ctx context.Context, id int) (entity.Segment, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.Segment{}, err
	}

	var (
		segment entity.Segment
		query   = fmt.Sprintf("SELECT * FROM %s WHERE id = $1",
			collectionSegments)
	)

	err = tx.QueryRowContext(ctx, query, id).
		Scan(&segment.ID, &segment.Name, &segment.AssignPercent)
	if err != nil {
		_ = tx.Rollback()
		return entity.Segment{}, err
	}

	return segment, tx.Commit()
}

func (s *SegmentsRepository) GetActiveUsersBySegmentID(ctx context.Context, id int) ([]entity.User, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return nil, err
	}

	var (
		users []entity.User
		query = fmt.Sprintf(
			"SELECT id, login, registered_at FROM %s JOIN %s ON %s.user_id = id WHERE %s.segment_id = $1",
			collectionUsers, collectionRelations, collectionRelations, collectionRelations)
	)

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Login, &user.RegisteredAt)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		_ = tx.Commit()
		return nil, err
	}

	return users, tx.Commit()
}

func (s *SegmentsRepository) GetAll(ctx context.Context) ([]entity.Segment, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return nil, err
	}

	var (
		segments []entity.Segment
		query    = fmt.Sprintf("SELECT * FROM %s",
			collectionSegments)
	)

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var segment entity.Segment
		if err = rows.Scan(&segment.ID, &segment.Name, &segment.AssignPercent); err != nil {
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

func (s *SegmentsRepository) DeleteByName(ctx context.Context, name string) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}

	var queryDeleteSegment = fmt.Sprintf("DELETE FROM %s WHERE name = $1 RETURNING id, name", collectionSegments)

	_, err = tx.ExecContext(ctx, queryDeleteSegment, name)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *SegmentsRepository) DeleteByID(ctx context.Context, id int) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}

	var query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", collectionSegments)

	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
