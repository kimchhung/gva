package utils

import "strings"

func AppendQuery(acc, query string) string {
	if acc == "" {
		return query
	}

	return acc + " AND " + query
}

func ColumnName(column string) string {
	if strings.Contains(column, ".") {
		parts := strings.Split(column, ".")
		for i, part := range parts {
			parts[i] = "`" + part + "`"
		}
		column = strings.Join(parts, ".")
	} else {
		column = "`" + column + "`"
	}

	return column
}
