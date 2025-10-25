package coretype

import "context"

type AppRouter interface {
	Register(ctx context.Context)
}
