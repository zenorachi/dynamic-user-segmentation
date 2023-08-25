package repository

import (
	"database/sql"
)

type RelationsRepository struct {
	db *sql.DB
}

func NewRelations(db *sql.DB) *RelationsRepository {
	return &RelationsRepository{db: db}
}

//func (r *RelationsRepository) CreateBySegmentID(ctx context.Context, relation entity.Relation) (int, error) {
//	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  false,
//	})
//	if err != nil {
//		return 0, err
//	}
//
//	var (
//		operationId          int
//		segmentName          string
//		queryGetSegmentName  = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", collectionSegments)
//		queryCreateRelation  = fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2)", collectionRelations)
//		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
//			collectionOperations)
//	)
//
//	err = tx.QueryRowContext(ctx, queryGetSegmentName, relation.SegmentID).Scan(&segmentName)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	_, err = tx.ExecContext(ctx, queryCreateRelation, relation.UserID, relation.SegmentID)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	err = tx.QueryRowContext(ctx, queryInsertOperation, relation.UserID, segmentName, entity.TypeAdd).Scan(&operationId)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	return operationId, tx.Commit()
//}
//
//func (r *RelationsRepository) CreateBySegmentName(ctx context.Context, userId int, segmentName string) (int, error) {
//	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  false,
//	})
//	if err != nil {
//		return 0, err
//	}
//
//	var (
//		segmentId            int
//		operationId          int
//		queryGetId           = fmt.Sprintf("SELECT id FROM %s WHERE name = $1", collectionSegments)
//		queryCreateRelation  = fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2)", collectionRelations)
//		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
//			collectionOperations)
//	)
//
//	err = tx.QueryRowContext(ctx, queryGetId, segmentName).Scan(&segmentId)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	_, err = tx.ExecContext(ctx, queryCreateRelation, userId, segmentId)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	err = tx.QueryRowContext(ctx, queryInsertOperation, userId, segmentName, entity.TypeAdd).Scan(&operationId)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	return operationId, tx.Commit()
//}
//
//func (r *RelationsRepository) GetByUserID(ctx context.Context, userId int) (entity.Relation, error) {
//	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  true,
//	})
//	if err != nil {
//		return entity.Relation{}, err
//	}
//
//	var (
//		relation entity.Relation
//		query    = fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", collectionRelations)
//	)
//
//	err = tx.QueryRowContext(ctx, query, userId).Scan(&relation.UserID, &relation.SegmentID)
//	if err != nil {
//		_ = tx.Rollback()
//		return entity.Relation{}, err
//	}
//
//	return relation, tx.Commit()
//}
//
//func (r *RelationsRepository) GetBySegmentID(ctx context.Context, segmentId int) (entity.Relation, error) {
//	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  true,
//	})
//	if err != nil {
//		return entity.Relation{}, err
//	}
//
//	var (
//		relation entity.Relation
//		query    = fmt.Sprintf("SELECT * FROM %s WHERE segment_id = $1", collectionRelations)
//	)
//
//	err = tx.QueryRowContext(ctx, query, segmentId).Scan(&relation.UserID, &relation.SegmentID)
//	if err != nil {
//		_ = tx.Rollback()
//		return entity.Relation{}, err
//	}
//
//	return relation, tx.Commit()
//}
//
//func (r *RelationsRepository) DeleteBySegmentID(ctx context.Context, relation entity.Relation) (int, error) {
//	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  false,
//	})
//	if err != nil {
//		return 0, err
//	}
//
//	var (
//		operationId          int
//		segmentName          string
//		queryGetSegmentName  = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", collectionSegments)
//		queryDeleteRelation  = fmt.Sprintf("DELETE FROM %s WHERE segment_id = $1", collectionRelations)
//		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
//			collectionOperations)
//	)
//
//	err = tx.QueryRowContext(ctx, queryGetSegmentName, relation.SegmentID).Scan(&segmentName)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	_, err = tx.ExecContext(ctx, queryDeleteRelation, relation.SegmentID)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	err = tx.QueryRowContext(ctx, queryInsertOperation, relation.UserID, segmentName, entity.TypeAdd).Scan(&operationId)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	return operationId, tx.Commit()
//}
//
//func (r *RelationsRepository) DeleteBySegmentName(ctx context.Context, userId int, segmentName string) (int, error) {
//	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  false,
//	})
//	if err != nil {
//		return 0, err
//	}
//
//	var (
//		segmentId            int
//		operationId          int
//		queryGetSegmentId    = fmt.Sprintf("SELECT id FROM %s WHERE name = $1", collectionSegments)
//		queryDeleteRelation  = fmt.Sprintf("DELETE FROM %s WHERE segment_id = $1", collectionRelations)
//		queryInsertOperation = fmt.Sprintf("INSERT INTO %s (user_id, segment_name, type) VALUES ($1, $2, $3) RETURNING id",
//			collectionOperations)
//	)
//
//	err = tx.QueryRowContext(ctx, queryGetSegmentId, segmentName).Scan(&segmentId)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	_, err = tx.ExecContext(ctx, queryDeleteRelation, segmentId)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	err = tx.QueryRowContext(ctx, queryInsertOperation, userId, segmentName, entity.TypeAdd).Scan(&operationId)
//	if err != nil {
//		_ = tx.Rollback()
//		return 0, err
//	}
//
//	return operationId, tx.Commit()
//}
