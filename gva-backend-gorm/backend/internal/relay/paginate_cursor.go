package relay

import (
	"backend/internal/relay/cursor"

	"gorm.io/gorm"
)

func setCursor(db *gorm.DB, direction string, orderBy map[string]any, opt *PaginateGlobalConfig) (*gorm.DB, error) {

	var queries []string
	var args []any
	var err error

	if len(orderBy) == 0 {
		orderBy = map[string]any{
			opt.PrimaryKey: "ASC",
		}
	}

	if direction == "after" {
		queries, args, err = cursor.LoadCursor(cursor.OpAfter, opt.encoder)(opt.Table, opt.Tables, opt.After, orderBy, opt.PrimaryKey, opt.inSensitiveColumn)
	} else {
		queries, args, err = cursor.LoadCursor(cursor.OpBefore, opt.encoder)(opt.Table, opt.Tables, opt.Before, orderBy, opt.PrimaryKey, opt.inSensitiveColumn)
	}

	if err != nil {
		return db, err
	}

	for i := range queries {
		db = db.Where(queries[i], args[i])
	}

	return db, err
}
