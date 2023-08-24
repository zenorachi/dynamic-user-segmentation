package entity

import "time"

type User struct {
	ID           int       `json:"id,omitempty"`
	Login        string    `json:"login"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registered_at,omitempty"`
}

func NewUser(login, email, password string) User {
	return User{
		Login:    login,
		Email:    email,
		Password: password,
	}
}
