package gormrepo

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ErrorOrNotExist struct {
	ErrCause error
	NotExist bool
}

func NewErrorOrNotExist(errCause error) *ErrorOrNotExist {
	return &ErrorOrNotExist{
		ErrCause: errCause,
		NotExist: errors.Is(errCause, gorm.ErrRecordNotFound),
	}
}
