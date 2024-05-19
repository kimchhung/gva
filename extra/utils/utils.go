package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	if key {
		return nil
	}

	return func(c *fiber.Ctx) bool { return true }
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
