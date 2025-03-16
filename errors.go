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
	if errors.Is(errCause, gorm.ErrRecordNotFound) {
		return &ErrorOrNotExist{
			ErrCause: errCause,
			NotExist: true,
		}
	}
	return &ErrorOrNotExist{
		ErrCause: errCause,
		NotExist: false,
	}
}
