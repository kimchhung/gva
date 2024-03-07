package app_err

import (
	"errors"
	"fmt"
)

type Error struct {
	HttpCode  int
	ErrorCode int
	Message   string
}

type Option func(*Error)

func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

func WithMessageFunc(fn func(previous string) string) Option {
	return func(err *Error) {
		err.Message = fn(err.Error())
	}
}

// overwrite original message
func WithMessage(msg string) Option {
	return func(err *Error) {
		err.Message = msg
	}
}

func Join(err error) Option {
	return func(pErr *Error) {
		pErr.Message = errors.Join(pErr, err).Error()
	}
}

func NewError(err *Error, opts ...Option) *Error {
	nErr := new(Error)
	*nErr = *err

	for _, opt := range opts {
		opt(nErr)
	}

	return nErr
}
