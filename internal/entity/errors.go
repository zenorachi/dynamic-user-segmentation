package entity

import "errors"

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrEmptyAuthHeader   = errors.New("empty authorization header")
	ErrInvalidAuthHeader = errors.New("invalid authorization header")
	ErrUserAlreadyExists = errors.New("user with such login already exist")
	ErrUserDoesNotExist  = errors.New("user does not exist")
	ErrIncorrectPassword = errors.New("incorrect password")
)
