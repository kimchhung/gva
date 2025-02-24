package relayt

type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

func NewErrRelay(message string) *Error {
	return &Error{message: message}
}
