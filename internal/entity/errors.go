package entity

import "errors"

var (
	ErrInvalidInput          = errors.New("invalid input")
	ErrEmptyAuthHeader       = errors.New("empty authorization header")
	ErrInvalidAuthHeader     = errors.New("invalid authorization header")
	ErrUserAlreadyExists     = errors.New("user with such login already exist")
	ErrUserDoesNotExist      = errors.New("user does not exist")
	ErrIncorrectPassword     = errors.New("incorrect password")
	ErrSessionDoesNotExist   = errors.New("session does not exist")
	ErrSegmentAlreadyExists  = errors.New("segment already exists")
	ErrSegmentDoesNotExist   = errors.New("segment/segments does not exist")
	ErrRelationAlreadyExists = errors.New("relation already exists")
	ErrRelationDoesNotExist  = errors.New("relation does not exist")
	ErrInvalidTTL            = errors.New("invalid ttl")
	ErrInvalidAssignPercent  = errors.New("assign percent must be greater than zero")
	//ErrTTLIsNotDefined       = errors.New("ttl is not defined")
)
