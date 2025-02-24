package gormq

// gorm query

import (
	"fmt"
	"backend/utils/json"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
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
	OpBetween      OpName = "between"
	OpBetweenEq    OpName = "betweenEq"
)

func Equal(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s = ?", column), value)
	}
}

func EqualFold(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("LOWER(%s) = LOWER(?)", column), value)
	}
}
func NotEqual(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s != ?", column), value)
	}
}

func In[T any](column string, value []T) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s IN (?)", column), value)
	}
}

func NotIn[T any](column string, value []T) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s NOT IN (?)", column), value)
	}
}

func Contains(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s LIKE ?", column), "%"+value.(string)+"%")
	}
}

func ContainsFold(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("LOWER(%s) LIKE ?", column), "%"+value.(string)+"%")
	}
}

func Gt(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s > ?", column), value)
	}
}

func Gte(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s >= ?", column), value)
	}
}

func Lt(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s < ?", column), value)
	}
}

func Lte(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s <= ?", column), value)
	}
}

func HasPrefix(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s LIKE ?", column), value.(string)+"%")
	}
}

func HasSuffix(column string, value any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s LIKE ?", column), "%"+value.(string))
	}
}

func BetweenEq[T any](column string, value []T) Option {
	return func(tx *gorm.DB) *gorm.DB {
		tx = Gte(column, value[0])(tx)
		tx = Lte(column, value[1])(tx)
		return tx
	}
}

func Between[T any](column string, value []T) Option {
	return func(tx *gorm.DB) *gorm.DB {
		tx = Gt(column, value[0])(tx)
		tx = Lt(column, value[1])(tx)
		return tx
	}
}

func IsNull(column string, value bool) Option {
	return func(tx *gorm.DB) *gorm.DB {
		if value {
			return tx.Where(fmt.Sprintf("%s IS NULL", column))
		} else {
			return tx.Where(fmt.Sprintf("%s IS NOT NULL", column))
		}
	}
}

func Where(opts ...Option) Option {
	return func(tx *gorm.DB) *gorm.DB {
		for _, opt := range opts {
			opt(tx)
		}
		return tx
	}
}

func WhereExpr(query any, args ...any) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query, args...)
	}
}

// sort=["id desc","created desc"]
func WithSorts(sorts []string, mapField map[string]string) Option {
	return func(tx *gorm.DB) *gorm.DB {
		for _, s := range sorts {
			column := strings.TrimSpace(s)

			if column == "" {
				continue
			}

			direction := "asc"

			if column[0] == '-' {
				direction = "desc"
				column = s[1:]
			}

			if dbcolumn, ok := mapField[column]; ok {
				column = dbcolumn
			} else {
				continue
			}
			tx = OrderBy(column, direction == "desc")(tx)
		}

		return tx
	}
}

type MapOption func(column string) string

// myUser => my_user
func ToSnake() MapOption {
	return strcase.ToSnake
}

func WithPrefix(tableName string, opts ...MapOption) MapOption {
	return func(column string) string {
		for _, opt := range opts {
			column = opt(column)
		}
		return tableName + "." + column
	}
}

func ToSnakeWithTablePrefix(tableName string) MapOption {
	return func(column string) string {
		return fmt.Sprintf("%s.%s", tableName, strcase.ToSnake(column))
	}
}

// myUser => my_users.user
func ToColumn(column string) MapOption {
	return func(_ string) string {
		return column
	}
}

// ToJSONColumn("data","title") => JSON_VALUE(data, '$.title')
func ToJSONColumn(column string, field string) MapOption {
	return func(_ string) string {
		return fmt.Sprintf("JSON_VALUE(%s, '$.%s')", column, field)
	}
}

func Ignore() MapOption {
	return func(column string) string {
		return column
	}
}

/*
store map column to db column

	ColumnMap = {
		"createdAt" : "posts.created_at",
		"user.createdAt" : "User.created_at"
	}
*/
type ColumnMap map[string]string

// Returns all column names in the map.
func (kv ColumnMap) Keys() []string {
	var keys []string
	for k := range kv {
		keys = append(keys, k)
	}
	return keys
}

// Returns all values in the map.
func (kv ColumnMap) Values() []string {
	var values []string
	for _, v := range kv {
		values = append(values, v)
	}
	return values
}

// Returns a new ColumnMap containing only the columns specified in the keys parameter,
// including only those columns that have a non-empty value in the original map.
func (kv ColumnMap) Pick(keys ...string) ColumnMap {
	pickMap := make(ColumnMap)
	for _, key := range keys {
		if val, exists := kv[key]; exists && val != "" {
			pickMap[key] = val
		}
	}
	return pickMap
}

// Returns a new ColumnMap without the columns specified in the keys parameter
func (kv ColumnMap) Omit(keys ...string) ColumnMap {
	omitMap := make(ColumnMap, 0)
	for _, key := range keys {
		if val, exists := kv[key]; !exists && val != "" {
			omitMap[key] = val
		}
	}
	return omitMap
}

func MapTableColumn(mapper map[string]MapOption) ColumnMap {
	result := make(map[string]string)

	for k, fn := range mapper {
		if fn == nil {
			result[k] = k
			continue
		}

		result[k] = fn(k)
	}

	return result
}

// keyword1,keyword1
func WithSearch(str string, columns ...string) Option {
	texts := strings.Split(str, ",")

	return func(tx *gorm.DB) *gorm.DB {
		andBracket := tx.Session(&gorm.Session{NewDB: true})
		if str == "" || len(texts) == 0 {
			return tx
		}

		for _, text := range texts {
			text = strings.TrimSpace(text)
			safeText := safeTextForSql(text)
			orBracket := andBracket.Session(&gorm.Session{NewDB: true})

			for _, column := range columns {
				safeColumn := safeColumn(column)
				orBracket = orBracket.Or(fmt.Sprintf("%s LIKE ?", safeColumn), safeText)
			}

			andBracket = andBracket.Where(orBracket)
		}

		return tx.Where(andBracket)
	}
}

// safeTextForSql sanitizes the input text by replacing unsafe characters for SQL queries.
func safeTextForSql(text string) string {
	// Replace wildcard asterisk (*) with SQL's LIKE wildcard (%)
	text = strings.ReplaceAll(text, "*", "%")

	// Escape single quotes for SQL queries
	text = strings.ReplaceAll(text, "'", "''")

	// Remove potentially harmful characters (e.g., semicolon, hyphen)
	re := regexp.MustCompile(`[;\-]`)
	text = re.ReplaceAllString(text, "")

	return text
}

// safeColumn sanitizes a column name, ensuring no harmful characters or syntax and wraps it in backticks.
func safeColumn(column string) string {
	var columns []string

	if strings.Contains(column, "JSON_VALUE") {
		return column
	}

	if strings.Contains(column, ".") {
		columns = strings.Split(column, ".")
	} else {
		columns = []string{column}
	}

	for i, t := range columns {
		// Remove harmful characters from each part of the column (e.g., semicolon, hyphen)
		re := regexp.MustCompile(`[;\-]`)
		t = re.ReplaceAllString(t, "")

		// Wrap column parts in backticks to prevent SQL injection
		columns[i] = fmt.Sprintf("`%s`", t)
	}

	return strings.Join(columns, ".")
}

func WithFilters(where map[string]map[OpName]any, mapField map[string]string) Option {
	return func(tx *gorm.DB) *gorm.DB {
		for column, filters := range where {
			if dbcolumn, ok := mapField[column]; ok {
				column = dbcolumn
			} else {
				continue
			}

			for operator, value := range filters {
				switch operator {
				case OpEqual:
					tx = Equal(column, value)(tx)
				case OpEqualFold:
					tx = EqualFold(column, value)(tx)
				case OpNotEqual:
					tx = NotEqual(column, value)(tx)
				case OpIn:
					list := make([]any, 0)
					if err := json.MustJSON(value).Out(&list); err == nil && len(list) > 0 {
						tx = In(column, list)(tx)
					}

				case OpNotIn:
					list := make([]any, 0)
					if err := json.MustJSON(value).Out(&list); err == nil && len(list) > 0 {
						tx = NotIn(column, list)(tx)
					}

				case OpContains:
					tx = Contains(column, value)(tx)
				case OpContainsFold:
					tx = ContainsFold(column, value)(tx)
				case OpGt:
					tx = Gt(column, value)(tx)
				case OpGte:
					tx = Gte(column, value)(tx)
				case OpLt:
					tx = Lt(column, value)(tx)
				case OpLte:
					tx = Lte(column, value)(tx)
				case OpHasPrefix:
					tx = HasPrefix(column, value)(tx)
				case OpHasSuffix:
					tx = HasSuffix(column, value)(tx)
				case OpIsNull:
					if v, ok := value.(bool); ok {
						tx = IsNull(column, v)(tx)
					}
				case OpBetween:
					list := make([]any, 0)
					if err := json.MustJSON(value).Out(&list); err == nil && len(list) == 2 {
						tx = Between(column, list)(tx)
					}

				case OpBetweenEq:
					list := make([]any, 0)
					if err := json.MustJSON(value).Out(&list); err == nil && len(list) == 2 {
						tx = BetweenEq(column, list)(tx)
					}
				}
			}
		}

		return tx
	}
}
