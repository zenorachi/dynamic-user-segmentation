package entity

import "time"

type User struct {
	ID           int       `json:"id,omitempty"`
	Login        string    `json:"login,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	RegisteredAt time.Time `json:"registered_at,omitempty"`
}
