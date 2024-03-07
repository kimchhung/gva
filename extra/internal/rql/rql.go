package rql

import (
	"bytes"
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"
)

//go:generate easyjson -omit_empty -disallow_unknown_fields -snake_case rql.go

// Query is the decoded result of the user input.
//
//easyjson:json
type Query struct {
	// Limit must be > 0 and <= to `LimitMaxValue`.
	Limit int `json:"limit,omitempty"`
	// Offset must be >= 0.
	Offset int `json:"offset,omitempty"`
	// Select contains the list of expressions define the value for the `SELECT` clause.
	// For example:
	//
	//	params, err := p.Parse([]byte(`{
	//		"select": ["name", "age"]
	//	}`))
	//
	Select []string `json:"select,omitempty"`
	// Select contains the list of expressions define the value for the `UPDATE` clause.
	// For example:
	//
	//	params, err := p.Parse([]byte(`{
	//		"update": ["name", "age"]
	//	}`))
	//
	Update []string `json:"update,omitempty"`
	// Sort contains list of expressions define the value for the `ORDER BY` clause.
	// In order to return the rows in descending order you can prefix your field with `-`.
	// For example:
	//
	//	params, err := p.Parse([]byte(`{
	//		"sort": ["name", "-age", "+redundant"]
	//	}`))
	//
	Sort []string `json:"sort,omitempty"`
	// Filter is the query object for building the value for the `WHERE` clause.
	// The full documentation of the supported operators is writtern in the README.
	// An example for filter object:
	//
	//	params, err := p.Parse([]byte(`{
	//		"filter": {
	//			"account": { "$like": "%github%" },
	//			"$or": [
	//				{ "city": "TLV" },
	//				{ "city": "NYC" }
	//			]
	//		}
	//	}`))
	//
	Filter map[string]interface{} `json:"filter,omitempty"`

	// ## in development
	// Aggregate is the query object for building the value for the `SELECT` clause when grouped.
	// An example for filter object:
	//
	//	params, err := p.Parse([]byte(`{
	//		"aggregate": {
	//			"gold": { "$sum": "gold_fieldname" },		// returns SUM(gold_fieldname) AS gold
	//			"silver": { "$avg": "silver_fieldname" },	// returns AVG(silver_fieldname) AS silver
	//			"bronze": { "$min": "bronze_fieldname" },	// returns MIN(bronze_fieldname) AS bronze
	//			"bronze": { "$max": "bronze_fieldname" },	// returns MAX(bronze_fieldname) AS bronze
	//			"bronze": { "$count": "bronze_fieldname" },	// returns COUNT(bronze_fieldname) AS bronze
	//		}
	//	}`))
	//
	Aggregate map[string]interface{} `json:"aggregate,omitempty"`

	// ## in development
	// Group is the query object for building the value for the `GROUP` clause and will be appended to the `SELECT` clause.
	// An example for filter object:
	//
	//	params, err := p.Parse([]byte(`{
	//		"group": ["name", "age"]
	//	}`))
	//
	Group []string `json:"group,omitempty"`
}

// Params is the parser output after calling to `Parse`. You should pass its
// field values to your query tool. For example, Suppose you use `gorm`:
//
//	params, err := p.Parse(b)
//	if err != nil {
//		return nil, err
//	}
//	var users []User
//	err := db.Where(params.FilterExp, params.FilterArgs...).
//		Order(params.Sort).
//		Find(&users).
//		Error
//	if err != nil {
//		return nil, err
//	}
//	return users, nil
type Params struct {
	// Limit represents the number of rows returned by the SELECT statement.
	Limit int
	// Offset specifies the offset of the first row to return. Useful for pagination.
	Offset int
	// Select contains the expression for the `SELECT` clause defined in the Query. If group is not empty, values are automatically replaced by the group string and the aggregate string.
	Select []string
	// Aggregate contains additional aggregated `SELECT` expressions that can be optionally appended to the select statement.
	Aggregate []string
	// Update contains the expression for the `UPDATE` clause defined in the Query.
	Update []string
	// Sort used as a parameter for the `ORDER BY` clause. For example, "age desc, name".
	Sort []string
	// FilterExp and FilterArgs come together and used as a parameters for the `WHERE` clause.
	//
	// examples:
	// 	1. Exp: "name = ?"
	//	   Args: "a8m"
	//
	//	2. Exp: "name = ? AND age >= ?"
	// 	   Args: "a8m", 22
	FilterExp  ExpString
	FilterArgs []interface{}
	// GroupBy contains the expression for the `GROUP BY` clause defined in the Query. Values are automatically added to select string.
	Group []string
}

type ExpString string

// ParseError is type of error returned when there is a parsing problem.
type ParseError struct {
	msg string
}

func (p ParseError) Error() string {
	return p.msg
}

// field is a configuration of a struct field.
type field struct {
	// Name of the column.
	Name string
	// Has a "sort" option in the tag.
	Sortable bool
	// Has a "filter" option in the tag.
	SortableCaseInsensitive bool
	// Has a "filter" option in the tag.
	Filterable bool
	// Has a "group" option in the tag.
	Groupable bool
	// Has a "aggregate" option in the tag.
	Aggregateable bool
	// Has a "update" option in the tag.
	Updateable bool
	// All supported operators for this field.
	FilterOps map[string]bool
	// ALl supported modifiers for this field.
	ModifierOps map[string]bool
	// Validation for the type. for example, unit8 greater than or equal to 0.
	ValidateFn func(interface{}) error
	// ConvertFn converts the given value to the type value.
	CovertFn func(interface{}) interface{}
}

// A Parser parses various types. The result from the Parse method is a Param object.
// It is safe for concurrent use by multiple goroutines except for configuration changes.
type Parser struct {
	Config
	fields map[string]*field
}

// NewParser creates a new Parser. it fails if the configuration is invalid.
func NewParser(c Config) (*Parser, error) {
	if err := c.defaults(); err != nil {
		return nil, err
	}
	p := &Parser{
		Config: c,
		fields: make(map[string]*field),
	}
	if err := p.init(); err != nil {
		return nil, err
	}
	return p, nil
}

// MustNewParser is like NewParser but panics if the configuration is invalid.
// It simplifies safe initialization of global variables holding a resource parser.
func MustNewParser(c Config) *Parser {
	p, err := NewParser(c)
	if err != nil {
		panic(err)
	}

	return p
}

// Parse parses the given buffer into a Param object. It returns an error
// if the JSON is invalid, or its values don't follow the schema of rql.
func (p *Parser) Parse(b []byte) (pr *Params, err error) {
	q := &Query{}
	if err := q.UnmarshalJSON(b); err != nil {
		return nil, &ParseError{"decoding buffer to *Query: " + err.Error()}
	}

	r, err := p.ParseQuery(q)
	if err != nil {
		return nil, err
	}

	// row group queries

	return r, nil
}

func (e ExpString) String() string {
	return string(e)
}

// Deprecated: use the following concept if you have a postgres driver
//
// s.Where(
// 	sql.P(func(b *sql.Builder) {
// 		split := strings.Split(string(params.FilterExp), "?")
// 		if len(split) != len(params.FilterArgs)+1 {
// 			panic("invalid number of args")
// 		}

//		for i, v := range params.FilterArgs {
//			b.WriteString(split[i])
//			b.Arg(v)
//		}
//		b.WriteString(split[len(split)-1])
//	}),
//
// )
func (e ExpString) PostgresString(offset int) string {
	// adapt for postgres
	// newExp := e.String()
	// cnt := strings.Count(newExp, "?")
	// for i := 1; i <= cnt; i++ {
	// 	newExp = strings.Replace(newExp, "?", fmt.Sprintf("$%v", i+offset), 1)
	// }
	// return newExp

	panic("deprecated")
}

// ParseQuery parses the given struct into a Param object. It returns an error
// if one of the query values don't follow the schema of rql.
func (p *Parser) ParseQuery(q *Query) (pr *Params, err error) {
	defer func() {
		if e := recover(); e != nil {
			perr, ok := e.(*ParseError)
			if !ok {
				panic(e)
			}
			err = perr
			pr = nil
		}
	}()
	pr = &Params{
		Limit: p.DefaultLimit,
	}
	expect(q.Offset >= 0, "offset must be greater than or equal to 0")
	pr.Offset = q.Offset
	if q.Limit != 0 {
		expect(q.Limit > 0 && q.Limit <= p.LimitMaxValue, "limit must be greater than 0 and less than or equal to %d", p.LimitMaxValue)
		pr.Limit = q.Limit
	}
	ps := p.newParseState()
	ps.and(q.Filter)
	pr.FilterExp = ExpString(ps.String())
	pr.FilterArgs = ps.values
	pr.Group = p.group(q.Group)
	pr.Select = p.sel(q.Select)

	aps := p.newParseState()
	aps.aggregate(q.Aggregate)
	agg := strings.Split(aps.String(), ", ")
	for _, a := range agg {
		if a != "" {
			pr.Aggregate = append(pr.Aggregate, a)
		}
	}

	pr.Sort = p.sort(q.Sort)
	if len(pr.Sort) == 0 && len(p.DefaultSort) > 0 && len(pr.Group) == 0 {
		pr.Sort = p.sort(p.DefaultSort)
	}
	pr.Update = p.validateUpdateColumnNames(q.Update)
	parseStatePool.Put(ps)
	return
}

// Column is the default function that converts field name into a database column.
// It used to convert the struct fields into their database names. For example:
//
//	Username => username
//	FullName => full_name
//	HTTPCode => http_code
func Column(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// put '.' if it is not a start or end of a word, current letter is an uppercase letter,
		// and previous letter is a lowercase letter (cases like: "UserName"), or next letter is
		// also a lowercase letter and previous letter is not "_".
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) &&
			(unicode.IsLower(rune(s[i-1])) ||
				unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1]))) {
			b.WriteString("_")
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

// init initializes the parser parsing state. it scans the fields
// in a breath-first-search order and for each one of the field calls parseField.
func (p *Parser) init() error {
	t := indirect(reflect.TypeOf(p.Model))
	l := list.New()
	for i := 0; i < t.NumField(); i++ {
		l.PushFront(t.Field(i))
	}
	for l.Len() > 0 {
		f := l.Remove(l.Front()).(reflect.StructField)
		_, ok := f.Tag.Lookup(p.TagName)
		switch t := indirect(f.Type); {
		// no matter what the type of this field. if it has a tag,
		// it is probably a filterable or sortable.
		case ok:
			if err := p.parseField(f); err != nil {
				return err
			}
		case t.Kind() == reflect.Struct:
			for i := 0; i < t.NumField(); i++ {
				f1 := t.Field(i)
				if !f.Anonymous {
					f1.Name = f.Name + p.FieldSep + f1.Name
				}
				if f.Name != "Edges" {
					l.PushFront(f1)
				}
			}
			// allow field without tag to be selected.
			if err := p.parseField(f); err != nil {
				return err
			}
		case f.Anonymous:
			p.Log("ignore embedded field %q that is not struct type", f.Name)
		default:
			// allow field without tag to be selected.
			if err := p.parseField(f); err != nil {
				return err
			}

		}
	}
	return nil
}

// parseField parses the given struct field tag, and add a rule
// in the parser according to its type and the options that were set on the tag.
func (p *Parser) parseField(sf reflect.StructField) error {
	f := &field{
		Name:        p.ColumnFn(sf.Name),
		CovertFn:    valueFn,
		FilterOps:   make(map[string]bool),
		ModifierOps: make(map[string]bool),
	}
	layout := time.RFC3339
	opts := strings.Split(sf.Tag.Get(p.TagName), ",")
	for _, opt := range opts {
		switch s := strings.TrimSpace(opt); {
		case s == "sort":
			f.Sortable = true
		case s == "filter":
			f.Filterable = true
		case s == "group":
			f.Groupable = true
		case s == "aggregate":
			f.Aggregateable = true
		case s == "update":
			f.Updateable = true
		case strings.HasPrefix(opt, "column"):
			f.Name = strings.TrimPrefix(opt, "column=")
		case strings.HasPrefix(opt, "layout"):
			layout = strings.TrimPrefix(opt, "layout=")
			// if it's one of the standard layouts, like: RFC822 or Kitchen.
			if ly, ok := layouts[layout]; ok {
				layout = ly
			}
			// test the layout on a value (on itself). however, some layouts are invalid
			// time values for time.Parse, due to formats such as _ for space padding and
			// Z for zone information.
			v := strings.NewReplacer("_", " ", "Z", "+").Replace(layout)
			if _, err := time.Parse(layout, v); err != nil {
				return fmt.Errorf("rql: layout %q is not parsable: %v", layout, err)
			}
		case s == "-" || s == "ignore":
			return nil
		case s == "":
			// when tag is missing allow select
		default:
			p.Log("Ignoring unknown option %q in struct tag", opt)
		}
	}
	var filterOps []Op
	var modifierOps []Op
	switch typ := indirect(sf.Type); typ.Kind() {
	case reflect.Bool:
		f.ValidateFn = validateBool
		filterOps = append(filterOps, EQ, NEQ, IN, ISNULL, ISNOTNULL)
		modifierOps = append(modifierOps, COUNT)
	case reflect.Array, reflect.Slice:
		f.ValidateFn = validateString
		filterOps = append(filterOps, EQ, NEQ, ISNULL, ISNOTNULL)
		// modifierOps = append(modifierOps, EQ)
	case reflect.String:
		f.ValidateFn = validateString
		f.SortableCaseInsensitive = true
		filterOps = append(filterOps, EQ, NEQ, IN, LT, LTE, GT, GTE, LIKE, ILIKE, ISNULL, ISNOTNULL)
		modifierOps = append(modifierOps, MIN, MAX, COUNT)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f.ValidateFn = validateInt
		f.CovertFn = convertInt
		filterOps = append(filterOps, EQ, NEQ, IN, LT, LTE, GT, GTE, ISNULL, ISNOTNULL)
		modifierOps = append(modifierOps, MIN, MAX, COUNT, SUM, AVG, ROUND, BALANCE)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		f.ValidateFn = validateUInt
		f.CovertFn = convertInt
		filterOps = append(filterOps, EQ, NEQ, IN, LT, LTE, GT, GTE, ISNULL, ISNOTNULL)
		modifierOps = append(modifierOps, MIN, MAX, COUNT, SUM, AVG, ROUND, BALANCE)
	case reflect.Float32, reflect.Float64:
		f.ValidateFn = validateFloat
		filterOps = append(filterOps, EQ, NEQ, IN, LT, LTE, GT, GTE, ISNULL, ISNOTNULL)
		modifierOps = append(modifierOps, MIN, MAX, COUNT, SUM, AVG, ROUND, BALANCE)
	case reflect.Struct:
		switch v := reflect.Zero(typ); v.Interface().(type) {
		case sql.NullBool:
			f.ValidateFn = validateBool
			filterOps = append(filterOps, EQ, NEQ, IN, ISNULL, ISNOTNULL)
			modifierOps = append(modifierOps, COUNT)
		case sql.NullByte:
			f.ValidateFn = validateString
			filterOps = append(filterOps, EQ, NEQ, ISNULL, ISNOTNULL)
			// modifierOps = append(modifierOps, EQ)
		case sql.NullString:
			f.ValidateFn = validateString
			f.SortableCaseInsensitive = true
			filterOps = append(filterOps, EQ, NEQ, IN, ISNULL, ISNOTNULL)
			modifierOps = append(modifierOps, MIN, MAX, COUNT)
		case sql.NullInt64:
			f.ValidateFn = validateInt
			f.CovertFn = convertInt
			filterOps = append(filterOps, EQ, NEQ, IN, LT, LTE, GT, GTE, AVG, ROUND, ISNULL, ISNOTNULL)
			modifierOps = append(modifierOps, MIN, MAX, COUNT, SUM, BALANCE)
		case sql.NullFloat64:
			f.ValidateFn = validateFloat
			filterOps = append(filterOps, EQ, NEQ, IN, LT, LTE, GT, GTE, AVG, ROUND, ISNULL, ISNOTNULL)
			modifierOps = append(modifierOps, MIN, MAX, COUNT, SUM, BALANCE)
		case time.Time:
			f.ValidateFn = validateTime(layout)
			f.CovertFn = convertTime(layout)
			filterOps = append(filterOps, EQ, NEQ, IN, LT, LTE, GT, GTE, ISNULL, ISNOTNULL)
			modifierOps = append(modifierOps, MIN, MAX, COUNT, TRUNC, EXTRACT)
		default:
			if !v.Type().ConvertibleTo(reflect.TypeOf(time.Time{})) {
				if !p.Config.DoNotLog {
					fmt.Printf("rql: field type '%v' for %q is not supported - only allowed for select\n", typ.Kind(), sf.Name)
				}
			}
			f.ValidateFn = validateTime(layout)
			f.CovertFn = convertTime(layout)
			filterOps = append(filterOps, EQ, NEQ, IN, LT, LTE, GT, GTE, ISNULL, ISNOTNULL)
			modifierOps = append(modifierOps, EQ)
		}
	default:
		if !p.Config.DoNotLog {
			fmt.Printf("rql: field type '%v' for %q is not supported - only allowed for select\n", typ.Kind(), sf.Name)
		}
	}
	for _, op := range filterOps {
		f.FilterOps[p.op(op)] = true
	}
	for _, op := range modifierOps {
		f.ModifierOps[op.String()] = true
	}
	p.fields[f.Name] = f
	return nil
}

type parseState struct {
	*Parser                     // reference of the parser config
	*bytes.Buffer               // query builder
	values        []interface{} // query values
}

var parseStatePool sync.Pool

func (p *Parser) newParseState() (ps *parseState) {
	if v := parseStatePool.Get(); v != nil {
		ps = v.(*parseState)
		ps.Reset()
		ps.values = nil
	} else {
		ps = new(parseState)
		// currently we're using an arbitrary size as the capacity of initial buffer.
		// What we can do in the future is to track the size of parse results, and use
		// the average value. Same thing applies to the `values` field below.
		ps.Buffer = bytes.NewBuffer(make([]byte, 0, 64))
	}
	ps.values = make([]interface{}, 0, 8)
	ps.Parser = p
	return
}

func (p *Parser) validateUpdateColumnNames(fields []string) []string {
	for _, field := range fields {
		expect(p.fields[field] != nil, "unrecognized field %q for select", field)
		expect(p.fields[field].Updateable, "update on field %q not allowed", field)
	}
	return fields
}

// sort build the sort clause.
func (p *Parser) sort(fields []string) []string {
	sortParams := make([]string, len(fields))
	for i, field := range fields {
		expect(field != "", "sort field can not be empty")
		var orderBy string
		// if the sort field prefixed by an order indicator.
		if order, ok := sortDirection[field[0]]; ok {
			orderBy = order
			field = field[1:]
		}
		expect(p.fields[field] != nil, "unrecognized key %q for sorting", field)
		expect(p.fields[field].Sortable, "field %q is not sortable", field)
		colName := ""
		if p.fields[field].SortableCaseInsensitive {
			colName = fmt.Sprintf("lower(%v)", p.colName(field))
		} else {
			colName = fmt.Sprintf("%v", p.colName(field))
		}
		if orderBy != "" {
			colName += " " + orderBy
		}
		sortParams[i] = colName
	}
	return sortParams
}

func (p *Parser) group(fields []string) []string {
	groupParams := make([]string, len(fields))
	for i, field := range fields {
		_, finalCol := p.applyModifiers(field, "group")
		groupParams[i] = finalCol
	}
	return groupParams
}

func (p *Parser) sel(fields []string) []string {
	selectFields := make([]string, len(fields))
	for i, field := range fields {
		fieldName, finalCol := p.applyModifiers(field, "select")
		if fieldName != finalCol {
			finalCol = fmt.Sprintf("%v AS %v", finalCol, fieldName)
		}
		selectFields[i] = finalCol
	}

	return selectFields
}

func (p *Parser) applyModifiers(field string, typ string) (fieldName string, res string) {
	split := strings.Split(field, "|")
	fieldName = split[0]
	res = fieldName
	expect(fieldName != "", "group field can not be empty")
	expect(p.fields[fieldName] != nil, "unrecognized key %q for applying modifiers", fieldName)
	expect(typ != "group" || p.fields[fieldName].Groupable, "field %q is not groupable", fieldName)

	if len(split) > 1 {
		if len(split) > 2 {
			panic("currencly only one modifer allowed")
		}

		for m := 1; m < len(split); m++ {
			res = p.applyOptions(res, split[m])
		}
	}
	return fieldName, res
}

func (p *Parser) applyOptions(val, cmd string) (res string) {
	split := strings.Split(cmd, ":")
	modifier := split[0]
	expect(modifier != "", "modifier command can not be empty")
	expect(p.fields[val] != nil, "unrecognized key %q for applying options", val)
	expect(p.fields[val].ModifierOps[modifier], "field %q has no modifier %v", val, modifier)

	options := []string{}
	if len(split) > 1 {
		options = split[1:]
	}
	return Op(modifier).FormatModifier(val, options)
}

func (p *parseState) aggregate(f map[string]interface{}) {
	// "gold": { "$sum": "gold_fieldname" },		// returns SUM(gold_fieldname) AS gold
	var i int
	for as, intrfc := range f {
		expect(validateCustomColumnString(as), "invalid datatype for aggregation field %q", as)
		if i > 0 {
			p.WriteString(", ")
		}
		agg, ok := intrfc.(map[string]interface{})
		if !ok {
			expect(false, "invalid datatype for aggregation field %q", as)
		} else {
			for k, v := range agg {
				col, ok := v.(string)
				if !ok {
					expect(false, "invalid datatype for aggregation field %q", as)
				}
				if _, ok := p.fields[col]; !ok {
					expect(false, "unrecognized field %q for aggregation", col)
				}
				expect(p.fields[col].Aggregateable, "field %q is not aggregateable", v)
				switch k {
				case p.op(COUNT):
					p.WriteString(fmt.Sprintf("COUNT(%v) AS %v", col, as))
				case p.op(SUM):
					p.WriteString(fmt.Sprintf("SUM(%v) AS %v", col, as))
				case p.op(AVG):
					p.WriteString(fmt.Sprintf("AVG(%v) AS %v", col, as))
				case p.op(MAX):
					p.WriteString(fmt.Sprintf("MAX(%v) AS %v", col, as))
				case p.op(MIN):
					p.WriteString(fmt.Sprintf("MIN(%v) AS %v", col, as))
				default:
					expect(false, "unrecognized key %q for aggregation", k)
				}
			}
		}
		i++
	}
}

func (p *parseState) and(f map[string]interface{}) {
	var i int
	for k, v := range f {
		if i > 0 {
			p.WriteString(" AND ")
		}
		switch {
		case k == p.op(OR):
			terms, ok := v.([]interface{})
			expect(ok, "$or must be type array")
			p.relOp(OR, terms)
		case k == p.op(AND):
			terms, ok := v.([]interface{})
			expect(ok, "$and must be type array")
			p.relOp(AND, terms)
		case k == p.op(NOT):
			terms, ok := v.([]interface{})
			expect(ok, "$not must be type array")
			p.notOp(NOT, terms)
		case p.fields[k] != nil:
			expect(p.fields[k].Filterable, "field %q is not filterable", k)
			p.field(p.fields[k], v)
		default:
			expect(false, "unrecognized key %q for filtering", k)
		}
		i++
	}
}

func (p *parseState) relOp(op Op, terms []interface{}) {
	var i int
	if len(terms) > 1 {
		p.WriteByte('(')
	}
	for _, t := range terms {
		if i > 0 {
			p.WriteByte(' ')
			p.WriteString(op.SQL())
			p.WriteByte(' ')
		}
		mt, ok := t.(map[string]interface{})
		expect(ok, "expressions for $%s operator must be type object", op)
		p.and(mt)
		i++
	}
	if len(terms) > 1 {
		p.WriteByte(')')
	}
}

func (p *parseState) notOp(op Op, terms []interface{}) {
	var i int
	p.WriteString(op.SQL())
	p.WriteByte(' ')
	if len(terms) > 1 {
		p.WriteByte('(')
	}
	for _, t := range terms {
		if i > 0 {
			p.WriteByte(' ')
			p.WriteString(AND.SQL())
			p.WriteByte(' ')
		}
		mt, ok := t.(map[string]interface{})
		expect(ok, "expressions for $%s operator must be type object", op)
		p.and(mt)
		i++
	}
	if len(terms) > 1 {
		p.WriteByte(')')
	}
}

func (p *parseState) field(f *field, v interface{}) {
	terms, ok := v.(map[string]interface{})
	// default equality check.
	if !ok {
		must(f.ValidateFn(v), "invalid datatype for field %q", f.Name)
		p.WriteString(p.fmtOp(f.Name, EQ))
		p.values = append(p.values, f.CovertFn(v))
	}
	var i int
	if len(terms) > 1 {
		p.WriteByte('(')
	}
	for opName, opVal := range terms {
		if i > 0 {
			p.WriteString(" AND ")
		}
		expect(f.FilterOps[opName], "can not apply op %q on field %q", opName, f.Name)
		if opName == p.op(ISNULL) || opName == p.op(ISNOTNULL) {
			p.WriteString(p.fmtOp(f.Name, Op(opName[1:])))
			p.values = append(p.values, nil)
		} else if opName == p.op(IN) {
			if valArr, ok := opVal.([]interface{}); !ok {
				expect(false, "invalid datatype for field %q", f.Name)
			} else {
				for _, inVal := range valArr {
					must(f.ValidateFn(inVal), "invalid datatype or format in array of field %q", f.Name)
					p.values = append(p.values, f.CovertFn(inVal))
				}
				p.WriteString(p.fmtOp(f.Name, Op(opName[1:]), len(valArr)))
			}
		} else {
			must(f.ValidateFn(opVal), "invalid datatype or format for field %q", f.Name)
			p.WriteString(p.fmtOp(f.Name, Op(opName[1:])))
			p.values = append(p.values, f.CovertFn(opVal))
		}
		i++
	}
	if len(terms) > 1 {
		p.WriteByte(')')
	}
}

// fmtOp create a string for the operation with a placeholder.
// for example: "name = ?", or "age >= ?".
func (p *Parser) fmtOp(field string, op Op, length ...int) string {
	colName := p.colName(field)
	if op == IN && len(length) > 0 && length[0] > 0 {
		args := make([]string, 0, length[0])
		for i := 0; i < length[0]; i++ {
			args = append(args, "?")
		}
		return colName + " " + op.SQL() + " (" + strings.Join(args, ",") + ")"
	}
	return colName + " " + op.SQL() + " ?"
}

// colName formats the query field to database column name in cases the user configured a custom
// field separator. for example: if the user configured the field separator to be ".", the fields
// like "address.name" will be changed to "address_name".
func (p *Parser) colName(field string) string {
	str := field
	if p.FieldSep != DefaultFieldSep {
		if p.Config.InterpretFieldSepAsNestedJsonbObject {
			split := strings.Split(field, p.FieldSep)
			str = split[0]
			for i := 1; i < len(split); i++ {
				if regexp.MustCompile(`^[0-9]+$`).MatchString(split[i]) {
					str += p.Config.JsonbSep + split[i]
				} else {
					str += p.Config.JsonbSep + "'" + split[i] + "'"
				}
			}

			i := strings.LastIndex(str, p.Config.JsonbSep)
			if i > 0 {
				str = str[:i] + strings.Replace(str[i:], p.Config.JsonbSep, p.Config.JsonbLastSep, 1)
			}
		} else {
			str = strings.Replace(field, p.FieldSep, DefaultFieldSep, -1)
		}
	}

	return str
}

func (p *Parser) op(op Op) string {
	return p.OpPrefix + string(op)
}

// expect panic if the condition is false.
func expect(cond bool, msg string, args ...interface{}) {
	if !cond {
		panic(&ParseError{fmt.Sprintf(msg, args...)})
	}
}

// must panics if the error is not nil.
func must(err error, msg string, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		panic(&ParseError{fmt.Sprintf(msg+": %s", args...)})
	}
}

// indirect returns the item at the end of indirection.
func indirect(t reflect.Type) reflect.Type {
	for ; t.Kind() == reflect.Ptr; t = t.Elem() {
	}
	return t
}

// --------------------------------------------------------
// Validators and Converters

func errorType(v interface{}, expected string) error {
	actual := "nil"
	if v != nil {
		actual = reflect.TypeOf(v).Kind().String()
	}
	return fmt.Errorf("expect <%s>, got <%s>", expected, actual)
}

// validate that the underlined element of given interface is a boolean.
func validateBool(v interface{}) error {
	if _, ok := v.(bool); !ok {
		return errorType(v, "bool")
	}
	return nil
}

// validate that the underlined element of given interface is a string.
func validateString(v interface{}) error {
	if _, ok := v.(string); !ok {
		return errorType(v, "string")
	}
	return nil
}

// validate that the underlined element of given interface is a float.
func validateFloat(v interface{}) error {
	if _, ok := v.(float64); !ok {
		return errorType(v, "float64")
	}
	return nil
}

// validate that the underlined element of given interface is an int.
func validateInt(v interface{}) error {
	n, ok := v.(float64)
	if !ok {
		return errorType(v, "int")
	}
	if math.Trunc(n) != n {
		return errors.New("not an integer")
	}
	return nil
}

// validate that the underlined element of given interface is an int and greater than 0.
func validateUInt(v interface{}) error {
	if err := validateInt(v); err != nil {
		return err
	}
	if v.(float64) < 0 {
		return errors.New("not an unsigned integer")
	}
	return nil
}

// validate that the underlined element of this interface is a "datetime" string.
func validateTime(layout string) func(interface{}) error {
	return func(v interface{}) error {
		s, ok := v.(string)
		if !ok {
			return errorType(v, "string")
		}
		_, err := time.Parse(layout, s)
		return err
	}
}

// convert float to int.
func convertInt(v interface{}) interface{} {
	return int(v.(float64))
}

// convert string to time object.
func convertTime(layout string) func(interface{}) interface{} {
	return func(v interface{}) interface{} {
		t, _ := time.Parse(layout, v.(string))
		return t
	}
}

// nop converter.
func valueFn(v interface{}) interface{} {
	return v
}

// layouts holds all standard time.Time layouts.
var layouts = map[string]string{
	"ANSIC":       time.ANSIC,
	"UnixDate":    time.UnixDate,
	"RubyDate":    time.RubyDate,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Kitchen":     time.Kitchen,
	"Stamp":       time.Stamp,
	"StampMilli":  time.StampMilli,
	"StampMicro":  time.StampMicro,
	"StampNano":   time.StampNano,
}

func validateCustomColumnString(columnName string) bool {
	re := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return re.MatchString(columnName)
}
