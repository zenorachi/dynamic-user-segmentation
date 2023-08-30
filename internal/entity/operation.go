package entity

import "time"

const (
	TypeAdd    = "added"
	TypeDelete = "deleted"
)

type Operation struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	SegmentName string    `json:"segment_name,omitempty"`
	Type        string    `json:"type,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}
