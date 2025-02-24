package gormq

// gorm query

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func OrderBy(field string, isDesc bool) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(clause.OrderByColumn{Column: clause.Column{Name: field}, Desc: isDesc})
	}
}
