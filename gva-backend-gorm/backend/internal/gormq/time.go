package gormq

// gorm query

import (
	"backend/internal/datetime"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	Option func(q *gorm.DB) *gorm.DB
	Tx     func(q *gorm.DB) error

	TimeBetweenValue struct {
		Start time.Time
		End   time.Time

		/* Bracket
		[] => gte,lte
		() => gt  lt
		(] => gt,lte
		[) => gte,lt
		*/
		Bracket string
	}

	TimeBetweenStr struct {
		Start string
		End   string
	}
)

func Empty(q *gorm.DB) *gorm.DB {
	return q
}

func WithDebug() Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Session(&gorm.Session{
			Logger: logger.Default,
		}).Debug()
	}
}

/*
WhereTimeBetween

	default value TimeBetweenValue{
	    Start 2024-07-03 00:00:00
	    End   2024-07-03 23:59:59:999
	    Bracket: "[]" -> gte,lte

		sql -> utc time -> WHERE date >= "2024-07-03 16:00:00" AND date <= "2024-07-07 15:59:59.999"
	}
*/
func WhereTimeBetween(field string, bets ...TimeBetweenValue) Option {
	return func(q *gorm.DB) *gorm.DB {
		var (
			session      = q
			isUseBracket = len(bets) > 1
		)

		if isUseBracket {
			session = q.Session(&gorm.Session{NewDB: true})
		}

		if len(bets) == 0 {
			bets = []TimeBetweenValue{
				{
					Start: datetime.Must().StartOf("day").ToTime(),
					End:   datetime.Must().EndOf("day").ToTime(),
				},
			}
		}

		for _, bet := range bets {
			brackets := []rune(bet.Bracket)
			if len(brackets) != 2 {
				brackets = []rune("[]")
			}

			switch brackets[0] {
			case '(':
				session = Gt(field, bet.Start)(session)
			case '[':
				session = Gte(field, bet.Start)(session)
			}

			switch brackets[1] {
			case ')':
				session = Lt(field, bet.End)(session)
			case ']':
				session = Lte(field, bet.End)(session)
			}
		}

		if isUseBracket {
			return q.Where(session)
		}

		return session
	}
}
