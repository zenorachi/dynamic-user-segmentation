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
		query = fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id",
			collectionSegments)
	)

	err = tx.QueryRowContext(ctx, query, segment.Name).Scan(&id)
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
