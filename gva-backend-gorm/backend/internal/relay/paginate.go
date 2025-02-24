package relay

import (
	"backend/internal/gormq"
	"backend/internal/relay/cursor"
	"backend/internal/relay/order"
	"backend/internal/relay/relaye"
	"backend/internal/relay/utils"
	"backend/internal/relay/where"

	"gorm.io/gorm"
)

var (
	DefaultEncoder cursor.Encoder
)

func init() {
	DefaultEncoder = relaye.NewEncoder("relay")
}

type PaginateConfig struct {
	First  *int    `json:"first,omitempty" query:"first"`
	Last   *int    `json:"last,omitempty" query:"last"`
	After  *string `json:"after,omitempty" query:"after"`
	Before *string `json:"before,omitempty" query:"before"`
}

type PaginateGlobalConfig struct {
	First  *int
	Last   *int
	After  *string
	Before *string
	Table  string

	/*
		Map field to use which table
			tables := map[string]string{
						"id":"posts"
						"myId":"users"
					}
			sql := "where posts.id and users.my_id...."
	*/
	Tables     *map[string]string
	PrimaryKey string

	where             any
	order             any
	inSensitiveColumn []string

	scopes                   []gormq.Option
	maxLimit                 *int
	isDisableCount           bool
	isCreateAllCursorForEdge bool

	// cursor encoder
	encoder cursor.Encoder
}

func NewPaginateGlobalConfig(options ...PaginateOption) *PaginateGlobalConfig {
	config := &PaginateGlobalConfig{
		PrimaryKey: "id",
		encoder:    DefaultEncoder,
	}

	for _, opt := range options {
		opt(config)
	}

	return config
}

type PaginateOption func(config *PaginateGlobalConfig)

// cursor
func Cursor(config *PaginateConfig) PaginateOption {
	return func(gConfig *PaginateGlobalConfig) {
		gConfig.First = config.First
		gConfig.Last = config.Last
		gConfig.Before = config.Before
		gConfig.After = config.After
	}
}

/*
Map field to use which table

	tables := map[string]string{
				"id":"posts"
				"created_at":"users"
			}

	sql := "where posts.id and users.created_at...."
*/
func Tables(tables *map[string]string) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.Tables = tables
	}
}

// ovewrite primary id to sth else
func PrimaryKey(key string) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.PrimaryKey = key
	}
}

// table name
func Table(table string) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.Table = table
	}
}

// add relay where
func Where(where any) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.where = where
	}
}

func WithEncoder(encoder cursor.Encoder) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.encoder = encoder
	}
}

func DisableCount() PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.isDisableCount = true
	}
}

func EnableCreateAllCursor() PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.isCreateAllCursorForEdge = true
	}
}

func WithCount(isEnable bool) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.isDisableCount = !isEnable
	}
}

// add relay order
func Order(order any) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.order = order
	}
}

// m = -1, will remove limit according to gorm.DB docs
func MaxLimit(n int) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.maxLimit = &n
	}
}

// extend gorm query
func Query(opts ...gormq.Option) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.scopes = opts
	}
}

/*
Map field to use inSensitiveColumn

	cuesor := map[string]string{
				"id":"1"
				"age":"3"
				"created_at":"2"
			}
	sorts="id,age,created_at"
	inSensitiveColumn="age"

	case after
	sql := "where id>1 and age>=3 and created_at>2"

	case before
	sql := "where id<1 and age<=3 and created_at<2"

Note: inSensitiveColumn require at least 2 fields in cursor
*/
func InSensitiveColumn(inSensitiveColumn ...string) PaginateOption {
	return func(config *PaginateGlobalConfig) {
		config.inSensitiveColumn = inSensitiveColumn
	}
}

func Paginate[Model any](db *gorm.DB, opts ...PaginateOption) (*Connection[Model], error) {
	option := NewPaginateGlobalConfig(opts...)

	if err := validationAndDecode(option); err != nil {
		return nil, err
	}

	for _, scope := range option.scopes {
		db = scope(db)
	}

	w, err := where.Do(db.Dialector.Name(), option.Table, option.Tables, option.where)
	if err != nil {
		return nil, err
	}

	stmt := where.Traverse(db, w)

	var totalCount int64
	if !option.isDisableCount {
		totalCount, err = getTotalCount[Model](db)
		if err != nil {
			return nil, err
		}
	}

	orderBy, err := utils.ConvertToMap(option.order)
	if err != nil {
		return nil, err
	}

	if option.After != nil {
		stmt, err = setCursor(stmt, "after", orderBy, option)
		if err != nil {
			return nil, err
		}
	}

	if option.Before != nil {
		stmt, err = setCursor(stmt, "before", orderBy, option)
		if err != nil {
			return nil, err
		}
	}

	orders, err := order.By(option.Table, option.Tables, option.order, option.Last != nil)
	if err != nil {
		return nil, err
	}

	stmt = order.ApplyOrder(stmt, orders)
	stmt = limit(stmt, option)

	var rows []*Model
	if err := stmt.Find(&rows).Error; err != nil {
		return nil, err
	}

	edges, err := convertToEdge(option.encoder, rows, utils.Keys(orderBy), option.PrimaryKey, option.isCreateAllCursorForEdge)
	if err != nil {
		return nil, err
	}

	pageInfo := &PageInfo{}
	_totalCount := int(totalCount)
	edgesLen := len(edges)
	pageInfo.SetHasPreviousPage(_totalCount, edgesLen, option)
	pageInfo.SetHasNextPage(_totalCount, edgesLen, option)

	if edgesLen > 0 {
		pageInfo.StartCursor = &edges[0].Cursor
		pageInfo.EndCursor = &edges[edgesLen-1].Cursor
	}

	return &Connection[Model]{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo:   pageInfo,
	}, nil
}

func PaginateTo[Model any](db *gorm.DB, result []*Model, opts ...PaginateOption) (*Connection[Model], error) {
	option := NewPaginateGlobalConfig(opts...)

	if err := validationAndDecode(option); err != nil {
		return nil, err
	}

	for _, scope := range option.scopes {
		db = scope(db)
	}

	w, err := where.Do(db.Dialector.Name(), option.Table, option.Tables, option.where)
	if err != nil {
		return nil, err
	}

	stmt := where.Traverse(db, w)

	var totalCount int64
	if !option.isDisableCount {
		totalCount, err = getTotalCount[Model](db)
		if err != nil {
			return nil, err
		}
	}

	orderBy, err := utils.ConvertToMap(option.order)
	if err != nil {
		return nil, err
	}

	if option.After != nil {
		stmt, err = setCursor(stmt, "after", orderBy, option)
		if err != nil {
			return nil, err
		}
	}

	if option.Before != nil {
		stmt, err = setCursor(stmt, "before", orderBy, option)
		if err != nil {
			return nil, err
		}
	}

	orders, err := order.By(option.Table, option.Tables, option.order, option.Last != nil)
	if err != nil {
		return nil, err
	}

	stmt = order.ApplyOrder(stmt, orders)
	stmt = limit(stmt, option)

	if err := stmt.Find(&result).Error; err != nil {
		return nil, err
	}

	edges, err := convertToEdge(option.encoder, result, utils.Keys(orderBy), option.PrimaryKey, option.isCreateAllCursorForEdge)
	if err != nil {
		return nil, err
	}

	pageInfo := &PageInfo{}
	_totalCount := int(totalCount)
	edgesLen := len(edges)
	pageInfo.SetHasPreviousPage(_totalCount, edgesLen, option)
	pageInfo.SetHasNextPage(_totalCount, edgesLen, option)

	if edgesLen > 0 {
		pageInfo.StartCursor = &edges[0].Cursor
		pageInfo.EndCursor = &edges[edgesLen-1].Cursor
	}

	return &Connection[Model]{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo:   pageInfo,
	}, nil
}
