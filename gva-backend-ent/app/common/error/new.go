package apperror

import (
	"strings"

	"github.com/gva/internal/lang"
)

type Error struct {
	HttpCode  int
	ErrorCode int
	Message   string

	subErrs            []error
	isDisableTranslate bool
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

func (e Error) IsDisableTranslate() bool {
	return e.isDisableTranslate
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

// Add prefix to original message
func Prefix(localeOpt lang.LocaleOption, prefix string) Option {
	return func(err *Error) {
		if localeOpt == nil {
			err.Message = prefix + err.Message
			return
		}

		translated := lang.T(localeOpt, err.Message)
		err.isDisableTranslate = true
		if !lang.Is(localeOpt, lang.LocaleZH) {
			prefix += " "
		}
		err.Message = prefix + translated
	}
}

// Add sufic to original message
func Suffix(localeOpt lang.LocaleOption, suffic string) Option {
	return func(err *Error) {
		if localeOpt == nil {
			err.Message = err.Message + suffic
			return
		}

		translated := lang.T(localeOpt, err.Message)
		err.isDisableTranslate = true
		if !lang.Is(localeOpt, lang.LocaleZH) {
			suffic = " " + suffic
		}
		err.Message = translated + suffic
	}
}

/*
this clone original error and apply modification change
by default error will be translated when response

	NewError(err,DisableTranslate()) // to disable translation when response
*/
func NewError(err *Error, opt Option, opts ...Option) *Error {
	nErr := new(Error)
	*nErr = *err
	opt(nErr)

	for _, op := range opts {
		op(nErr)
	}

	return nErr
}

// to disable translation
func DisableTranslate() Option {
	return func(err *Error) {
		err.isDisableTranslate = true
	}
}
