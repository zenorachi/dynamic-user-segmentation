package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

type OperationsRepository struct {
	db *sql.DB
}

func NewOperations(db *sql.DB) *OperationsRepository {
	return &OperationsRepository{db: db}
}

func (o *OperationsRepository) CreateBySegmentID(ctx context.Context, relation entity.Relation) (int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return 0, err
	}

	var (
		operationId          int
		segmentName          string
		queryGetSegmentName  = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", collectionSegments)
		queryCreateRelation  = fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2)", collectionRelations)
		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
			collectionOperations)
	)

	err = tx.QueryRowContext(ctx, queryGetSegmentName, relation.SegmentID).Scan(&segmentName)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_, err = tx.ExecContext(ctx, queryCreateRelation, relation.UserID, relation.SegmentID)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.QueryRowContext(ctx, queryInsertOperation, relation.UserID, segmentName, entity.TypeAdd).Scan(&operationId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return operationId, tx.Commit()
}

func (o *OperationsRepository) CreateBySegmentName(ctx context.Context, userId int, segmentName string) (int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return 0, err
	}

	var (
		segmentId            int
		operationId          int
		queryGetId           = fmt.Sprintf("SELECT id FROM %s WHERE name = $1", collectionSegments)
		queryCreateRelation  = fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2)", collectionRelations)
		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
			collectionOperations)
	)

	err = tx.QueryRowContext(ctx, queryGetId, segmentName).Scan(&segmentId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_, err = tx.ExecContext(ctx, queryCreateRelation, userId, segmentId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.QueryRowContext(ctx, queryInsertOperation, userId, segmentName, entity.TypeAdd).Scan(&operationId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return operationId, tx.Commit()
}

func (o *OperationsRepository) GetByUserID(ctx context.Context, userId int) (entity.Relation, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.Relation{}, err
	}

	var (
		relation entity.Relation
		query    = fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", collectionRelations)
	)

	err = tx.QueryRowContext(ctx, query, userId).Scan(&relation.UserID, &relation.SegmentID)
	if err != nil {
		_ = tx.Rollback()
		return entity.Relation{}, err
	}

	return relation, tx.Commit()
}

func (o *OperationsRepository) GetBySegmentID(ctx context.Context, segmentId int) (entity.Relation, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.Relation{}, err
	}

	var (
		relation entity.Relation
		query    = fmt.Sprintf("SELECT * FROM %s WHERE segment_id = $1", collectionRelations)
	)

	err = tx.QueryRowContext(ctx, query, segmentId).Scan(&relation.UserID, &relation.SegmentID)
	if err != nil {
		_ = tx.Rollback()
		return entity.Relation{}, err
	}

	return relation, tx.Commit()
}

func (o *OperationsRepository) GetByBothIDs(ctx context.Context, userId, segmentId int) (entity.Relation, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return entity.Relation{}, err
	}

	var (
		relation entity.Relation
		query    = fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND segment_id = $2",
			collectionRelations)
	)

	err = tx.QueryRowContext(ctx, query, userId, segmentId).Scan(&relation.UserID, &relation.SegmentID)
	if err != nil {
		_ = tx.Rollback()
		return entity.Relation{}, err
	}

	return relation, tx.Commit()
}

func (o *OperationsRepository) DeleteBySegmentID(ctx context.Context, relation entity.Relation) (int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return 0, err
	}

	var (
		operationId          int
		segmentName          string
		queryGetSegmentName  = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", collectionSegments)
		queryDeleteRelation  = fmt.Sprintf("DELETE FROM %s WHERE segment_id = $1", collectionRelations)
		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
			collectionOperations)
	)

	err = tx.QueryRowContext(ctx, queryGetSegmentName, relation.SegmentID).Scan(&segmentName)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_, err = tx.ExecContext(ctx, queryDeleteRelation, relation.SegmentID)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.QueryRowContext(ctx, queryInsertOperation, relation.UserID, segmentName, entity.TypeDelete).Scan(&operationId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return operationId, tx.Commit()
}

func (o *OperationsRepository) DeleteBySegmentName(ctx context.Context, userId int, segmentName string) (int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return 0, err
	}

	var (
		segmentId            int
		operationId          int
		queryGetSegmentId    = fmt.Sprintf("SELECT id FROM %s WHERE name = $1", collectionSegments)
		queryDeleteRelation  = fmt.Sprintf("DELETE FROM %s WHERE segment_id = $1", collectionRelations)
		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
			collectionOperations)
	)

	err = tx.QueryRowContext(ctx, queryGetSegmentId, segmentName).Scan(&segmentId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_, err = tx.ExecContext(ctx, queryDeleteRelation, segmentId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.QueryRowContext(ctx, queryInsertOperation, userId, segmentName, entity.TypeDelete).Scan(&operationId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return operationId, tx.Commit()
}

func (o *OperationsRepository) GetOperations(ctx context.Context, userIds ...int) ([]entity.Operation, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return nil, err
	}

	var operations []entity.Operation

	args, query := o.generateGetQuery(userIds...)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var operation entity.Operation
		err = rows.Scan(&operation.ID, &operation.UserID, &operation.SegmentName, &operation.Type, &operation.Date)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		operations = append(operations, operation)
	}

	return operations, tx.Commit()
}

func (o *OperationsRepository) generateGetQuery(userIds ...int) ([]any, string) {
	var (
		query string
		args  []interface{}
	)

	if len(userIds) > 0 {
		query = fmt.Sprintf("SELECT * FROM %s WHERE user_id IN (", collectionOperations)
		for i, id := range userIds {
			query += fmt.Sprintf("$%d", i+1)
			args = append(args, id)
			if i < len(userIds)-1 {
				query += ","
			}
		}
		query += ")"
	} else {
		query = fmt.Sprintf("SELECT * FROM %s", collectionOperations)
	}

	return args, query
}
