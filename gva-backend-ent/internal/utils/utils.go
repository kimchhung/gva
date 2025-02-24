package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func IsEnabled(key bool) func(c echo.Context) bool {
	if key {
		return nil
	}

	return func(c echo.Context) bool { return true }
}

func SetIfEmpty[T comparable](dest *T, value T) {
	if dest == nil {
		dest = new(T)
	}

	if IsEmpty(dest) {
		dest = &value
	}
}

func IsEmpty[T comparable](v T) bool {
	var zero T
	return v == zero
}

func PanicIfErr(prefix string, err error) {
	if err == nil {
		return
	}
	panic(fmt.Errorf("%s %v", prefix, err))
}

func Async[A any](f func() A) <-chan A {
	ch := make(chan A, 1)
	go func() {
		ch <- f()
	}()
	return ch
}
