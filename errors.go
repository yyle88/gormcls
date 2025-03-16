package gormrepo

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ErrorOrNotExist struct {
	errCause error
	notExist bool
}

func NewErrorOrNotExist(errCause error) *ErrorOrNotExist {
	if errors.Is(errCause, gorm.ErrRecordNotFound) {
		return &ErrorOrNotExist{
			errCause: errCause,
			notExist: true,
		}
	}
	return &ErrorOrNotExist{
		errCause: errCause,
		notExist: false,
	}
}

func (e *ErrorOrNotExist) ErrCause() error {
	return e.errCause
}

func (e *ErrorOrNotExist) NotExist() bool {
	return e.notExist
}
