package gormq

// gorm query

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func WithSelect(columns ...string) Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Select(columns)
	}
}

func WithOmit(columns ...string) Option {
	return func(q *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return q
		}

		return q.Omit(columns...)
	}
}

func WithHardDelete() Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Unscoped()
	}
}

func WithoutLimitAndOffset() Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Offset(-1).Limit(-1)
	}
}

func WithCount(count *int64, countOpts ...Option) Option {
	return func(q *gorm.DB) *gorm.DB {
		newSession := q.Session(&gorm.Session{Initialized: true})
		for _, opt := range countOpts {
			opt(newSession)
		}
		WithoutLimitAndOffset()(newSession).Count(count)
		return q
	}
}

func WithModel(model interface{}) Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Model(model)
	}
}

func WithLockUpdate() Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate})
	}
}

func WhereStruct[T comparable](model T) Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where(model)
	}
}

func Preload(name string, args ...interface{}) func(q *gorm.DB) *gorm.DB {
	return func(q *gorm.DB) *gorm.DB {
		return q.Preload(name, args)
	}
}

func WithPageAndLimit(page int, limit int) Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Offset((page - 1) * limit).Limit(limit)
	}
}

func If(condition bool, opt Option) Option {
	return func(q *gorm.DB) *gorm.DB {
		if condition {
			q = opt(q)
		}
		return q
	}
}

func Multi(opts ...Option) Option {
	return func(q *gorm.DB) *gorm.DB {
		for _, opt := range opts {
			q = opt(q)
		}
		return q
	}
}
