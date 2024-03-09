package app_err

import (
	"strings"
)

type Error struct {
	HttpCode  int
	ErrorCode int
	Message   string

	subErrs []error
}

type Option func(*Error)

const (
	seperator = ", "
)

func (e *Error) Error() string {
	msgs := make([]string, 0)
	msgs = append(msgs, e.Message)

	for _, subErr := range e.subErrs {
		if msg := subErr.Error(); msg != "" {
			msgs = append(msgs, msg)
		}
	}

	return strings.Join(msgs, seperator)
}

func (e *Error) Join(err error) *Error {
	e.subErrs = append(e.subErrs, err)
	return e
}

// overwrite original message with raw message
func Join(err error) Option {
	return func(_err *Error) {
		_err.Join(err)
	}
}

// overwrite original message with raw message
func Message(text string) Option {
	return func(err *Error) {
		err.Message = text
	}
}

// overwrite original message with raw message
func MessageFunc(fn func(message string) string) Option {
	return func(err *Error) {
		err.Message = fn(err.Message)
	}
}

// this clone original error and apply modification change
func NewError(err *Error, opt Option, opts ...Option) *Error {
	nErr := new(Error)
	*nErr = *err
	opt(nErr)

	for _, op := range opts {
		op(nErr)
	}

	return nErr
}
