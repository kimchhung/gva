package app_err

import (
	"context"
	"errors"
	"fmt"

	"github.com/kimchhung/gva/extra/lang"
)

type Error struct {
	HttpCode  int
	ErrorCode int
	Message   string

	isTranslate bool
}

type Option func(*Error)

func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

func (e *Error) Translate(ctx context.Context) {
	if e.isTranslate {
		return
	}

	e.Message = lang.T(e.Message, lang.WithContext(ctx))
	e.SetTranslated()
}

func (e *Error) SetTranslated() {
	e.isTranslate = true
}

func Join(ctx context.Context, err error) Option {
	return func(pErr *Error) {
		pErr.Translate(ctx)
		pErr.Message = errors.Join(pErr, err).Error()
	}
}

// overwrite original message
func WithMessage(msg string) Option {
	return func(err *Error) {
		err.Message = msg
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
