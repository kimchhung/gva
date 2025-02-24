package where

import (
	"fmt"
	"backend/internal/relay/utils"
	"strings"
)

type OpName string

const (
	OpEqual        OpName = "equal"
	OpEqualFold    OpName = "equalFold"
	OpNotEqual     OpName = "notEqual"
	OpIn           OpName = "in"
	OpNotIn        OpName = "notIn"
	OpContains     OpName = "contains"
	OpContainsFold OpName = "containsFold"
	OpGt           OpName = "gt"
	OpGte          OpName = "gte"
	OpLt           OpName = "lt"
	OpLte          OpName = "lte"
	OpHasPrefix    OpName = "hasPrefix"
	OpHasSuffix    OpName = "hasSuffix"
	OpIsNull       OpName = "isNull"
)

func filter(dialector, column string, input map[string]any) (query string, args []any) {

	column = utils.ColumnName(column)

	for key, value := range input {
		if strings.Contains((fmt.Sprintf("%T", value)), "map[string]") {
			return mapFilter(dialector, column, key, value.(map[string]any))
		}

		switch OpName(key) {
		case OpEqual:
			query = utils.AppendQuery(query, column+" = ?")
			args = append(args, value)
		case OpNotEqual:
			query = utils.AppendQuery(query, column+" != ?")
			args = append(args, value)
		case OpEqualFold:
			query = utils.AppendQuery(query, "LOWER("+column+") = LOWER(?)")
			args = append(args, value)
		case OpIn:
			query = utils.AppendQuery(query, column+" IN (?)")
			args = append(args, value)
		case OpNotIn:
			query = utils.AppendQuery(query, column+" NOT IN (?)")
			args = append(args, value)
		case OpContains:
			query = utils.AppendQuery(query, column+" LIKE ?")
			args = append(args, "%"+value.(string)+"%")
		case OpContainsFold:
			query = utils.AppendQuery(query, "LOWER("+column+") LIKE LOWER(?)")
			args = append(args, "%"+value.(string)+"%")
		case OpGt:
			query = utils.AppendQuery(query, column+" > ?")
			args = append(args, value)
		case OpGte:
			query = utils.AppendQuery(query, column+" >= ?")
			args = append(args, value)
		case OpLt:
			query = utils.AppendQuery(query, column+" < ?")
			args = append(args, value)
		case OpLte:
			query = utils.AppendQuery(query, column+" <= ?")
			args = append(args, value)
		case OpHasPrefix:
			query = utils.AppendQuery(query, column+" LIKE ?")
			args = append(args, value.(string)+"%")
		case OpHasSuffix:
			query = utils.AppendQuery(query, column+" LIKE ?")
			args = append(args, "%"+value.(string))
		case OpIsNull:
			if value.(bool) {
				query = utils.AppendQuery(query, column+" IS NULL")
			} else {
				query = utils.AppendQuery(query, column+" IS NOT NULL")
			}
		}
	}
	return
}
