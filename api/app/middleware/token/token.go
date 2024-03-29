// Simple token middleware as example

package token

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Next       func(c *fiber.Ctx) bool
	HeaderName string
	VerifyFunc func(c *fiber.Ctx, headerValue string) error
}

type Option func(*Config)
type JwtKey struct{}

func NewConfig(config *Config) Option {
	return func(_c *Config) {
		*_c = *config
	}
}

func WithHeaderName(headerName string) Option {
	return func(c *Config) {
		c.HeaderName = headerName
	}
}

func WithSkip(isSkip func(c *fiber.Ctx) bool) Option {
	return func(c *Config) {
		c.Next = isSkip
	}
}

func VerifyFunc(VerifyFunc func(c *fiber.Ctx, headerValue string) error) Option {
	return func(c *Config) {
		c.VerifyFunc = VerifyFunc
	}
}

// New creates a new middleware handler
func New(config Option, opts ...Option) fiber.Handler {
	// Set default config
	cfg := &Config{
		HeaderName: "authorization",
	}

	// Override default config
	for _, opt := range append(opts, config) {
		opt(cfg)
	}

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		if err := cfg.VerifyFunc(c, c.Get(cfg.HeaderName, "")); err != nil {
			return err
		}

		return c.Next()
	}
}
