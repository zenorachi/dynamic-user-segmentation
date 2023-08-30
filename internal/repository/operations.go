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

func (o *OperationsRepository) CreateRelationsBySegmentIDs(ctx context.Context, userId int, segmentIDs []int) ([]int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	var (
		operationsIDs        []int
		operationId          int
		segmentName          string
		queryGetSegmentName  = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", collectionSegments)
		queryCreateRelation  = fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2)", collectionRelations)
		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
			collectionOperations)
	)

	for _, segmentId := range segmentIDs {
		err = tx.QueryRowContext(ctx, queryGetSegmentName, segmentId).Scan(&segmentName)
		if err != nil {
			return nil, err
		}

		_, err = tx.ExecContext(ctx, queryCreateRelation, userId, segmentId)
		if err != nil {
			return nil, err
		}

		err = tx.QueryRowContext(ctx, queryInsertOperation, userId, segmentName, entity.TypeAdd).Scan(&operationId)
		if err != nil {
			return nil, err
		}

		operationsIDs = append(operationsIDs, operationId)
	}

	return operationsIDs, tx.Commit()
}

func (o *OperationsRepository) CreateRelationsBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	var (
		segmentId            int
		operationsIDs        []int
		operationId          int
		queryGetId           = fmt.Sprintf("SELECT id FROM %s WHERE name = $1", collectionSegments)
		queryCreateRelation  = fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2)", collectionRelations)
		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
			collectionOperations)
	)

	for _, segmentName := range segmentsNames {
		err = tx.QueryRowContext(ctx, queryGetId, segmentName).Scan(&segmentId)
		if err != nil {
			return nil, err
		}

		_, err = tx.ExecContext(ctx, queryCreateRelation, userId, segmentId)
		if err != nil {
			return nil, err
		}

		err = tx.QueryRowContext(ctx, queryInsertOperation, userId, segmentName, entity.TypeAdd).Scan(&operationId)
		if err != nil {
			return nil, err
		}

		operationsIDs = append(operationsIDs, operationId)
	}

	return operationsIDs, tx.Commit()
}

func (o *OperationsRepository) DeleteRelationsBySegmentIDs(ctx context.Context, userId int, segmentsIDs []int) ([]int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	var (
		operationsIDs        []int
		operationId          int
		segmentName          string
		queryGetSegmentName  = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", collectionSegments)
		queryDeleteRelation  = fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_id = $2", collectionRelations)
		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
			collectionOperations)
	)

	for _, segmentId := range segmentsIDs {
		err = tx.QueryRowContext(ctx, queryGetSegmentName, segmentId).Scan(&segmentName)
		if err != nil {
			return nil, err
		}

		result, err := tx.ExecContext(ctx, queryDeleteRelation, userId, segmentId)
		if err != nil {
			return nil, err
		}
		if deletedRows, _ := result.RowsAffected(); deletedRows == 0 {
			return nil, entity.ErrRelationDoesNotExist
		}

		err = tx.QueryRowContext(ctx, queryInsertOperation, userId, segmentName, entity.TypeDelete).Scan(&operationId)
		if err != nil {
			return nil, err
		}

		operationsIDs = append(operationsIDs, operationId)
	}

	return operationsIDs, tx.Commit()
}

func (o *OperationsRepository) DeleteRelationsBySegmentNames(ctx context.Context, userId int, segmentsNames []string) ([]int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	var (
		segmentId            int
		operationsIDs        []int
		operationId          int
		queryGetSegmentId    = fmt.Sprintf("SELECT id FROM %s WHERE name = $1", collectionSegments)
		queryDeleteRelation  = fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_id = $2", collectionRelations)
		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
			collectionOperations)
	)

	for _, segmentName := range segmentsNames {
		err = tx.QueryRowContext(ctx, queryGetSegmentId, segmentName).Scan(&segmentId)
		if err != nil {
			return nil, err
		}

		result, err := tx.ExecContext(ctx, queryDeleteRelation, userId, segmentId)
		if err != nil {
			return nil, err
		}
		if deletedRows, _ := result.RowsAffected(); deletedRows == 0 {
			return nil, entity.ErrRelationDoesNotExist
		}

		err = tx.QueryRowContext(ctx, queryInsertOperation, userId, segmentName, entity.TypeDelete).Scan(&operationId)
		if err != nil {
			return nil, err
		}

		operationsIDs = append(operationsIDs, operationId)
	}

	return operationsIDs, tx.Commit()
}

func (o *OperationsRepository) GetOperations(ctx context.Context, year, month int, userIDs ...int) ([]entity.Operation, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  true,
	})
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	var operations []entity.Operation

	args, query := o.generateGetOperationsQuery(year, month, userIDs...)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var operation entity.Operation
		err = rows.Scan(&operation.UserID, &operation.SegmentName, &operation.Type, &operation.Date)
		if err != nil {
			return nil, err
		}

		operations = append(operations, operation)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return operations, tx.Commit()
}

func (o *OperationsRepository) generateGetOperationsQuery(year, month int, userIDs ...int) ([]any, string) {
	var (
		query string
		args  []interface{}
		i     int
		id    int
	)

	if len(userIDs) > 0 {
		query = fmt.Sprintf("SELECT user_id, segment_name, type, date FROM %s WHERE user_id IN (", collectionOperations)
		for i, id = range userIDs {
			query += fmt.Sprintf("$%d", i+1)
			args = append(args, id)
			if i < len(userIDs)-1 {
				query += ","
			}
		}
		query += fmt.Sprintf(") AND EXTRACT(YEAR FROM date) = $%d AND EXTRACT(MONTH FROM date) = $%d", i+2, i+3)
	} else {
		query = fmt.Sprintf("SELECT user_id, segment_name, type, date FROM %s WHERE EXTRACT(YEAR FROM date) = $1 AND EXTRACT(MONTH FROM date) = $2", collectionOperations)
	}

	args = append(args, year, month)

	return args, query
}
