package entity

type Segment struct {
	ID            int     `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	AssignPercent float64 `json:"assign_percent,omitempty"`
}
