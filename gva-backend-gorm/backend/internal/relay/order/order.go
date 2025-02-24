package order

import (
	"fmt"
	"backend/internal/relay/utils"
	"strings"

	"gorm.io/gorm"
)

func By(table string, tables *map[string]string, input any, reverse bool) ([]string, error) {
	filter, err := utils.ConvertToMap(input)
	if err != nil {
		return nil, err
	}

	orderMap := make(map[string]string)
	for k, v := range filter {
		orderMap[k] = fmt.Sprintf("%v", v)
	}

	return traverse(table, tables, orderMap, reverse), nil
}

func ApplyOrder(stmt *gorm.DB, orders []string) *gorm.DB {
	for _, order := range orders {
		if order == "" {
			continue
		}
		parts := strings.Split(order, " ")
		order = utils.ColumnName(parts[0])
		if len(parts) > 1 {
			order += " " + parts[1]
		}
		stmt = stmt.Order(order)
	}
	return stmt
}
