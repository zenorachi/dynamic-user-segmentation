package entity

import "time"

const (
	TypeAdd    = "added"
	TypeDelete = "deleted"
)

type Operation struct {
	ID          int
	UserID      int
	SegmentName string
	Type        string
	Date        time.Time
}

func NewOperation(userId int, segmentName string, opType string) Operation {
	return Operation{
		UserID:      userId,
		SegmentName: segmentName,
		Type:        opType,
	}
}
