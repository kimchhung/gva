package cursor

import (
	"fmt"
	"backend/internal/relay/relayt"
	"backend/internal/relay/utils"
)

type CursorOperator string

const (
	OpAfter    CursorOperator = ">"
	OpBefore   CursorOperator = "<"
	OpAfterEq  CursorOperator = ">="
	OpBeforeEq CursorOperator = "<="
)

func (co CursorOperator) Reverse() CursorOperator {
	switch co {
	case OpAfter:
		return OpBefore
	case OpBefore:
		return OpAfter
	case OpAfterEq:
		return OpBeforeEq
	case OpBeforeEq:
		return OpAfterEq
	}

	return co
}

func (co CursorOperator) InSensitive() CursorOperator {
	switch co {
	case OpAfter:
		return OpAfterEq
	case OpBefore:
		return OpBeforeEq
	}

	return co
}

type CursorLoader func(table string, tables *map[string]string, cursorRaw *string, orderBy map[string]any, primaryKey string, inSensitiveColumn []string) (queries []string, args []any, err error)

func CursorMapWithTable(cursorMap map[string]any, table string, tables *map[string]string) map[string]any {
	if hasTable := tables != nil || table != ""; !hasTable {
		return cursorMap
	}

	cursorMapWithTable := make(map[string]any)
	for column, v := range cursorMap {
		tableNameWithColumn := utils.WithTable(column, table, tables)
		cursorMapWithTable[tableNameWithColumn] = v
	}

	return cursorMapWithTable
}

func LoadCursor(defaultInequality CursorOperator, encoder Encoder) CursorLoader {
	return func(table string, tables *map[string]string, cursorRaw *string, orderBy map[string]any, primaryKey string, inSensitiveColumn []string) (queries []string, args []any, err error) {
		if cursorRaw == nil {
			return
		}

		cursorMap, err := ParseCursor(encoder, *cursorRaw)
		if err != nil {
			return
		}

		cursorMap = CursorMapWithTable(cursorMap, table, tables)
		if len(orderBy) == 0 {
			orderBy = map[string]any{primaryKey: "ASC"}
		}

		mapOrder := make(map[string]any, len(orderBy))
		for field, value := range orderBy {
			_, column := utils.SplitOrderKey(field)
			column = utils.WithTable(column, table, tables)
			mapOrder[column] = value
		}

		sensitiveColumn := make(map[string]struct{})
		if len(orderBy) > 1 {
			for _, column := range inSensitiveColumn {
				column = utils.WithTable(column, table, tables)
				sensitiveColumn[column] = struct{}{}
			}
		}

		if len(orderBy) != len(cursorMap) {
			err = relayt.NewErrRelay("cursorMap and orderBy field mismatch")
			return
		}

		for field, value := range cursorMap {
			operator := defaultInequality
			direction, ok := mapOrder[field]
			if !ok {
				err = relayt.NewErrRelay(fmt.Sprintf("cursor and orderBy field mismatch: %s", field))
				return
			}

			if _, ok := sensitiveColumn[field]; ok {
				operator = operator.InSensitive()
			}

			queries = append(queries, fmt.Sprintf("%s %s ?", field, inequality(operator, fmt.Sprintf("%v", direction))))
			args = append(args, value)
		}
		return

	}
}

func inequality(operator CursorOperator, direction string) CursorOperator {
	if direction == "DESC" || direction == "desc" {
		return operator.Reverse()
	}

	return operator
}
