package entity

import "errors"

var (
	ErrInvalidInput          = errors.New("invalid input")
	ErrEmptyAuthHeader       = errors.New("empty authorization header")
	ErrInvalidAuthHeader     = errors.New("invalid authorization header")
	ErrUserAlreadyExists     = errors.New("user with such login/email already exists")
	ErrUserDoesNotExist      = errors.New("user does not exist")
	ErrIncorrectPassword     = errors.New("incorrect password")
	ErrSessionDoesNotExist   = errors.New("session does not exist")
	ErrSegmentAlreadyExists  = errors.New("segment already exists")
	ErrSegmentDoesNotExist   = errors.New("segment/segments does not exist")
	ErrRelationAlreadyExists = errors.New("relation already exists")
	ErrRelationDoesNotExist  = errors.New("relation does not exist")
	ErrInvalidTTL            = errors.New("invalid ttl")
	ErrInvalidAssignPercent  = errors.New("assign percent must be greater than zero")
	ErrInvalidHistoryPeriod  = errors.New("invalid history period")
	ErrGDriveIsNotAvailable  = errors.New("google drive service is not available")
	ErrFileNotFound          = errors.New("file not found")
	ErrPageIsOutOfBounds     = errors.New("page is out of bounds")
)
