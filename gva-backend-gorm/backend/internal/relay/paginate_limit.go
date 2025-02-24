package relay

import (
	"gorm.io/gorm"
)

func limit(db *gorm.DB, config *PaginateGlobalConfig) *gorm.DB {
	var limit int

	if config.maxLimit != nil {
		limit = *config.maxLimit
	} else {
		limit = 100
	}

	if config.First != nil && *config.First < limit {
		limit = *config.First
	}

	if config.Last != nil && *config.Last < limit {
		limit = *config.Last
	}

	return db.Limit(limit)
}
