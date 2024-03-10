package rql

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"slices"
	"strings"
	"unicode"
)

// Op is a filter operator used by rql.
type Op string

// SQL returns the SQL representation of the operator.
func (o Op) SQL() string {
	return opFormat[o]
}

func (o Op) FormatModifier(val string, options []string) string {
	if v, ok := modFormatter[o]; ok {
		return v(val, options)
	}
	return val
}

func (o Op) String() string {
	return string(o)
}

// Operators that support by rql.
const (
	EQ        = Op("eq")        // =
	NEQ       = Op("neq")       // <>
	LT        = Op("lt")        // <
	GT        = Op("gt")        // >
	LTE       = Op("lte")       // <=
	GTE       = Op("gte")       // >=
	IN        = Op("in")        // IN "PATTERN"
	LIKE      = Op("like")      // LIKE "PATTERN" (case sensitive LIKE)
	ILIKE     = Op("ilike")     // ILIKE "PATTERN" (case insensitive LIKE)
	ISNULL    = Op("isnull")    // IS NULL
	ISNOTNULL = Op("isnotnull") // IS NOT NULL
	NOT       = Op("not")       // disjunction
	OR        = Op("or")        // disjunction
	AND       = Op("and")       // conjunction
	COUNT     = Op("count")     // aggregation
	SUM       = Op("sum")       // aggregation
	AVG       = Op("avg")       // aggregation
	MIN       = Op("min")       // aggregation
	MAX       = Op("max")       // aggregation
	BALANCE   = Op("balance")   // aggregation
	TRUNC     = Op("trunc")     // aggregation
	EXTRACT   = Op("extract")   // aggregation
	ROUND     = Op("round")     // aggregation
)

// Default values for configuration.
const (
	DefaultTagName      = "rql"
	DefaultOpPrefix     = "$"
	DefaultFieldSep     = "_"
	DefaultJsonbSep     = "->"
	DefaultJsonbLastSep = "->>"
	DefaultLimit        = 25
	DefaultMaxLimit     = 100
	Offset              = "offset"
	Limit               = "limit"
)

var (

	// A sorting expression can be optionally prefixed with + or - to control the
	// sorting direction, ascending or descending. For example, '+field' or '-field'.
	// If the predicate is missing or empty then it defaults to '+'
	sortDirection = map[byte]string{
		'+': "asc",
		'-': "desc",
	}

	// date|trunc:month
	// date|trunc:day
	// date|trunc:year
	// date|round:0.01
	modFormatter = map[Op]func(val string, options []string) string{
		TRUNC: func(val string, options []string) string {
			expect(val != "", "trunc requires a value")
			expect(len(options) == 1, "trunc requires exactly one option")
			expect(slices.Contains([]string{"day", "month", "year"}, options[0]), fmt.Sprintf("trunc has no option %v", options[0]))
			return fmt.Sprintf("DATE_TRUNC('%v', %v)", options[0], val)
		},
		EXTRACT: func(val string, options []string) string {
			expect(val != "", "truncate requires a value")
			expect(len(options) == 1, "truncate requires exactly one option")
			expect(slices.Contains([]string{"day", "month", "year"}, options[0]), fmt.Sprintf("truncate has no option %v", options[0]))
			return fmt.Sprintf("EXTRACT(%v from %v)", options[0], val)
		},
		BALANCE: func(val string, options []string) string {
			expect(val != "", "balance requires a value")
			expect(len(options) == 1, "balance requires exactly one option")
			expect(strings.Split(options[0], "_")[0] == "acc", fmt.Sprintf("balance option must start with acc_ %v", options[0]))
			return fmt.Sprintf(`sum(
				case when debit_account_id = '%[1]v' then %[2]v else 0 end
				+ case when credit_account_id = '%[1]v' then %[2]v * -1 else 0 end
			)`, options[0], val)
		},
		ROUND: func(val string, options []string) string {
			expect(val != "", "round requires a value")
			if len(options) < 1 {
				options = append(options, "0")
			}
			return fmt.Sprintf("ROUND(%v, %v)", val, options[0])
		},
		MIN: func(val string, options []string) string {
			expect(val != "", "min requires a value")
			expect(len(options) == 0, "min accepts no options")
			return fmt.Sprintf("MIN(%v)", val)
		},
		MAX: func(val string, options []string) string {
			expect(val != "", "max requires a value")
			expect(len(options) == 0, "max accepts no options")
			return fmt.Sprintf("MAX(%v)", val)
		},
		SUM: func(val string, options []string) string {
			expect(val != "", "sum requires a value")
			expect(len(options) == 0, "sum accepts no options")
			return fmt.Sprintf("SUM(%v)", val)
		},
		COUNT: func(val string, options []string) string {
			expect(val != "", "max requires a value")
			expect(len(options) == 0, "max accepts no options")
			return fmt.Sprintf("COUNT(%v)", val)
		},
		AVG: func(val string, options []string) string {
			expect(val != "", "max requires a value")
			expect(options != nil || len(options) != 0, "max accepts no options")
			return fmt.Sprintf("AVG(%v)", val)
		},
	}

	opFormat = map[Op]string{
		EQ:        "=",
		NEQ:       "<>",
		LT:        "<",
		GT:        ">",
		LTE:       "<=",
		GTE:       ">=",
		IN:        "IN",
		LIKE:      "LIKE",
		ILIKE:     "ILIKE",
		ISNULL:    `IS`,
		ISNOTNULL: `IS NOT`,
		NOT:       "NOT",
		OR:        "OR",
		AND:       "AND",
	}
)

// Config is the configuration for the parser.
type Config struct {
	// TagName is an optional tag name for configuration. t defaults to "rql".
	TagName string
	// Model is the resource definition. The parser is configured based on the its definition.
	// For example, given the following struct definition:
	//
	//	type User struct {
	//		Age	 int	`rql:"filter,sort"`
	// 		Name string	`rql:"filter"`
	// 	}
	//
	// In order to create a parser for the given resource, you will do it like so:
	//
	//	var QueryParser = rql.MustNewParser(
	// 		Model: User{},
	// 	})
	//
	Model interface{}
	// OpPrefix is the prefix for operators. it defaults to "$". for example, in order
	// to use the "gt" (greater-than) operator, you need to prefix it with "$".
	// It similar to the MongoDB query language.
	OpPrefix string
	// FieldSep is the separator for nested fields in a struct. For example, given the following struct:
	//
	//	type User struct {
	// 		Name 	string	`rql:"filter"`
	//		Address	struct {
	//			City string `rql:"filter"``
	//		}
	// 	}
	//
	// We assume the schema for this struct contains a column named "address_city". Therefore, the default
	// separator is underscore ("_"). But, you can change it to "." for convenience or readability reasons.
	// Then you will be able to query your resource like this:
	//
	//	{
	//		"filter": {
	//			"address.city": "DC"
	// 		}
	//	}
	//
	// The parser will automatically convert it to underscore ("_"). If you want to control the name of
	// the column, use the "column" option in the struct definition. For example:
	//
	//	type User struct {
	// 		Name 	string	`rql:"filter,column=full_name"`
	// 	}
	//
	FieldSep string
	// JsonbSep replaces the field Separator with JsonbSep to access nested jsonb objects in postgres
	JsonbSep string
	// replaces the last jsonbSep with JsonbLastSep to access nested jsonb objects in postgres without quotes
	JsonbLastSep string
	// InterpretFieldSepAsNestedJsonbObject replaces the fieldSep with -> to access nested jsonb objects in postgres
	InterpretFieldSepAsNestedJsonbObject bool
	// ColumnFn is the function that translate the struct field string into a table column.
	// For example, given the following fields and their column names:
	//
	//	FullName => "full_name"
	// 	CreatedAt => "createdAt"
	//
	// It is preferred that you will follow the same convention that your ORM or other DB helper use.
	// For example, If you are using `gorm` you want to se this option like this:
	//
	//	var QueryParser = rql.MustNewParser(
	// 		ColumnFn: gorm.ToDBName,
	// 	})
	//
	ColumnFn func(string) string

	// ColumnFnDB is the function that translate the struct field string into a table column.
	// For example, given the following fields and their column names:
	//
	//	fullName => "full_name"
	// 	hTTPPort => "http_port"
	ColumnNameFn func(string) string

	// add table prefix to all column
	// createdAt => table.createdAt
	Table string

	// createdAt => users.createdAt
	MapColumnName map[string]string

	// Log the the logging function used to log debug information in the initialization of the parser.
	// It defaults `to log.Printf`.
	Log func(string, ...interface{})
	// DefaultLimit is the default value for the `Limit` field that returns when no limit supplied by the caller.
	// It defaults to 25.
	DefaultLimit int
	// LimitMaxValue is the upper boundary for the limit field. User will get an error if the given value is greater
	// than this value. It defaults to 100.
	LimitMaxValue int
	// DefaultSort is the default value for the 'Sort' field that returns when no sort expression is supplied by the caller.
	// It defaults to an empty string slice.
	DefaultSort []string

	// If true logs won't be printed to stdout
	DoNotLog bool
}

// defaults sets the default configuration of Config.
func (c *Config) defaults() error {
	if c.Model == nil {
		return errors.New("rql: 'Model' is a required field")
	}
	if indirect(reflect.TypeOf(c.Model)).Kind() != reflect.Struct {
		return errors.New("rql: 'Model' must be a struct type")
	}
	if c.Log == nil {
		c.Log = log.Printf
	}
	if c.ColumnFn == nil {
		c.ColumnFn = Column
	}
	if c.MapColumnName == nil {
		c.MapColumnName = make(map[string]string)
	}
	defaultString(&c.TagName, DefaultTagName)
	defaultString(&c.OpPrefix, DefaultOpPrefix)
	defaultString(&c.FieldSep, DefaultFieldSep)
	defaultString(&c.JsonbSep, DefaultJsonbSep)
	defaultString(&c.JsonbLastSep, DefaultJsonbLastSep)
	defaultInt(&c.DefaultLimit, DefaultLimit)
	defaultInt(&c.LimitMaxValue, DefaultMaxLimit)
	return nil
}

func defaultString(s *string, v string) {
	if *s == "" {
		*s = v
	}
}

func defaultInt(i *int, v int) {
	if *i == 0 {
		*i = v
	}
}

func PascalToCamelCase(s string) string {
	if !strings.Contains(s, "_") {
		// If the string does not contain underscores, it's already in camelCase.
		// Convert the first character to lowercase and return the original string.
		r := []rune(s)
		r[0] = unicode.ToLower(r[0])
		return string(r)
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return strings.ReplaceAll(string(r), "_", "")
}

// camelCaseToSnakeCase converts a camelCase string to snake_case.
func CamelCaseToSnakeCase(s string) string {
	var sb strings.Builder
	for i, r := range s {
		if i == 0 {
			sb.WriteRune(unicode.ToLower(r)) // Convert the first character to lowercase
		} else if unicode.IsUpper(r) {
			sb.WriteString("_")
			sb.WriteRune(unicode.ToLower(r)) // No need to convert to lowercase here
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
