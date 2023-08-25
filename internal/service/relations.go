package service

//
//import (
//	"context"
//	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
//	"github.com/zenorachi/dynamic-user-segmentation/internal/repository"
//	"time"
//)
//
//type RelationsTTLService struct {
//	repo repository.RelationsTTL
//}
//
//func NewRelationsTTL(repo repository.RelationsTTL) *RelationsTTLService {
//	return &RelationsTTLService{repo: repo}
//}
//
//func (r *RelationsTTLService) GetTTL(ttl string) (time.Duration, error) {
//	if ttl != "" {
//		ttlDuration, err := time.ParseDuration(ttl)
//		if err != nil {
//			return 0, err
//		}
//		return ttlDuration, nil
//	}
//
//	return 0, entity.ErrTTLIsNotDefined
//}
//
//func (r *RelationsTTLService) Create(ctx context.Context, relationsTTLs []entity.RelationTTL) error {
//	return r.repo.Create(ctx, relationsTTLs)
//}
//
//func (r *RelationsTTLService) DeleteAfterTTLBySegmentID(ctx context.Context) error {
//	return r.repo.DeleteAfterTTLBySegmentID(ctx)
//}
//
//func (r *RelationsTTLService) ScheduleCleanup(ctx context.Context) {
//	ticker := time.NewTicker(5 * time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			_ = r.DeleteAfterTTLBySegmentID(ctx)
//		case <-ctx.Done():
//			return
//		}
//	}
//}
