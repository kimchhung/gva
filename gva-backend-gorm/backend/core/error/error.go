package coreerror

import (
	"fmt"
	"strings"
)

var (
	errorByCode = map[int]*Error{}
)

// Informational responses (100 – 199)
// Successful responses (200 – 299)
// Redirection messages (300 – 399)
// Client error responses (400 – 499)
// Server error responses (500 – 599)
func newError(httpCode int, errorCode int, message string) *Error {
	if _, exists := errorByCode[errorCode]; exists {
		// If the error code already exists, return an error.
		panic(fmt.Sprintf("[ErrorByCode]: error code %d is not unique", errorCode))
	}

	err := &Error{
		HttpCode:  httpCode,
		ErrorCode: errorCode,
		Message:   message,
	}

	errorByCode[errorCode] = err
	return err
}

func NewPublicError(httpCode int, errorCode int, message string) *Error {
	return &Error{
		HttpCode:  httpCode,
		ErrorCode: errorCode,
		Message:   message,
		isPublic:  true,
	}
}

type Error struct {
	HttpCode  int    `json:"httpCode"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`

	subErrs []*Error

	isTranslated bool
	isPublic     bool
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

// same as update but always new instance
func (e *Error) Copy(updators ...Option) *Error {
	return NewError(e, updators...)
}

func (e *Error) IsTranslated() bool {
	return e.isTranslated
}

func (e *Error) IsPublic() bool {
	return e.isPublic
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

func Join(err *Error) Option {
	return func(_err *Error) {
		_err.subErrs = append(_err.subErrs, err)
	}
}

// overwrite original message with raw message
func Message(text string) Option {
	return func(err *Error) {
		err.Message = text
	}
}

func AppendMessage(text string) Option {
	return func(err *Error) {
		err.Message += text
	}
}

// overwrite original message with raw message
func MessageFunc(fn func(prev string) string) Option {
	return func(err *Error) {
		err.Message = fn(err.Message)
	}
}

func Translate(fn func(prev string) string) Option {
	return func(err *Error) {
		if err.isTranslated {
			return
		}
		err.Message = fn(err.Message)
		err.isTranslated = true
	}
}

func NewError(err *Error, opts ...Option) *Error {
	nErr := new(Error)
	*nErr = *err

	for _, op := range opts {
		op(nErr)
	}

	return nErr
}

func DisableTranslate() Option {
	return func(err *Error) {
		err.isTranslated = true
	}
}
