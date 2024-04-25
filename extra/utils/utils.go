package utils

import (
	"github.com/gofiber/fiber/v2"
)

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	if key {
		return nil
	}

	return func(c *fiber.Ctx) bool { return true }
}

func SetIfEmpty[T comparable](dest T, value T) {
	if IsEmpty(dest) {
		dest = value
	}
}

// set value if original value is nil
func SetIfNil[T comparable](dest *T, value T) {
	if dest == nil {
		v := new(T)
		*v = value
	}
}

func IsEmpty[T comparable](v T) bool {
	var zero T
	return v == zero
}
