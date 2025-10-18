package apperror

import (
	"strings"
)

type Error struct {
	HttpCode  int    `json:"httpCode"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`

	subErrs            []error
	isDisableTranslate bool
	isPublic           bool
}

type Option func(*Error)

const (
	seperator = ", "
)

func (e *Error) Update(updators ...Option) *Error {
	for _, updator := range updators {
		updator(e)
	}
	return e
}

func (e *Error) IsPublic() bool {
	return e.isPublic
}

// same as update but always new instance
func (e *Error) Copy(updators ...Option) *Error {
	return NewError(e, updators...)
}

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

func (e Error) IsDisableTranslate() bool {
	return e.isDisableTranslate
}

// overwrite original message with raw message
func Join(err error) Option {
	return func(_err *Error) {
		_err.subErrs = append(_err.subErrs, err)
	}
}

// overwrite original message with raw message
func WithMessage(text string) Option {
	return func(err *Error) {
		err.Message = text
	}
}

// overwrite original message with raw message
func WithMessageFunc(fn func(prev string) string) Option {
	return func(err *Error) {
		err.Message = fn(err.Message)
	}
}

func WithTranslator(fn func(prev string) string) Option {
	return func(err *Error) {
		if err.isDisableTranslate {
			return
		}
		err.Message = fn(err.Message)
		err.isDisableTranslate = true
	}
}

/*
this clone original error and apply modification change
by default error will be translated when response

	NewError(err,DisableTranslate()) // to disable translation when response
*/
func NewError(err *Error, opts ...Option) *Error {
	nErr := new(Error)
	*nErr = *err

	for _, op := range opts {
		op(nErr)
	}

	return nErr
}

func WithDisableTranslate() Option {
	return func(err *Error) {
		err.isDisableTranslate = true
	}
}
