package entity

import "time"

type Session struct {
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

func NewSession(refreshToken string, expiresAt time.Time) Session {
	return Session{
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}
}
