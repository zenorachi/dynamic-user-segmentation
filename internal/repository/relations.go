package repository

//
//import (
//	"context"
//	"database/sql"
//	"fmt"
//	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
//)
//
//type RelationTTLRepository struct {
//	db *sql.DB
//}
//
//func NewRelationsTTL(db *sql.DB) *RelationTTLRepository {
//	return &RelationTTLRepository{db: db}
//}
//
//func (r *RelationTTLRepository) Create(ctx context.Context, relationsTTLs []entity.RelationTTL) error {
//	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  false,
//	})
//	if err != nil {
//		return err
//	}
//
//	query := fmt.Sprintf("INSERT INTO %s (user_id, segment_id, expires_at) VALUES ($1, $2, $3)", collectionRelationsTTL)
//
//	for _, relationTTL := range relationsTTLs {
//		_, err = tx.ExecContext(ctx, query, relationTTL.UserID, relationTTL.SegmentID, relationTTL.ExpiresAt)
//		if err != nil {
//			_ = tx.Rollback()
//			return err
//		}
//	}
//
//	return tx.Commit()
//}
//
//func (r *RelationTTLRepository) DeleteAfterTTLBySegmentID(ctx context.Context) error {
//	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
//		Isolation: sql.LevelSerializable,
//		ReadOnly:  false,
//	})
//	if err != nil {
//		return err
//	}
//
//	query := fmt.Sprintf(
//		"WITH deleted_relations AS ("+
//			"DELETE FROM %s "+
//			"WHERE (user_id, segment_id) IN (SELECT user_id, segment_id FROM %s WHERE expires_at <= NOW()) RETURNING user_id, segment_id) "+
//			"INSERT INTO %s (user_id, segment_name, type) "+
//			"SELECT dr.user_id, s.name, 'deleted' "+
//			"FROM deleted_relations dr "+
//			"JOIN segments s ON dr.segment_id = s.id "+
//			"JOIN users u ON dr.user_id = u.id;",
//		collectionRelations, collectionRelationsTTL, collectionOperations)
//
//	_, err = tx.ExecContext(ctx, query)
//	if err != nil {
//		_ = tx.Rollback()
//		return err
//	}
//
//	return tx.Commit()
//}
