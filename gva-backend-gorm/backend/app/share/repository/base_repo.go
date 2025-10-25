package repository

import (
	"context"
	"errors"
	"backend/internal/dbtx"
	"backend/internal/gormq"
	"backend/internal/relay"

	"gorm.io/gorm"
)

type IBaseRepository[T any] interface {
	Create(ctx context.Context, record *T) (*T, error)
	CreateMany(ctx context.Context, record []T) ([]T, error)
	Delete(ctx context.Context, options ...gormq.Option) error
	DeleteById(ctx context.Context, id uint, options ...gormq.Option) error
	GetById(ctx context.Context, id uint, options ...gormq.Option) (*T, error)
	GetMany(ctx context.Context, options ...gormq.Option) ([]T, error)
	GetManyTo(ctx context.Context, result interface{}, options ...gormq.Option) error
	GetManyAndCount(ctx context.Context, resp any, count *int64, options ...gormq.Option) error
	GetOne(ctx context.Context, options ...gormq.Option) (*T, error)
	GetOneOrCreate(ctx context.Context, record *T, option gormq.Option, options ...gormq.Option) (bool, *T, error)
	GetOneTo(ctx context.Context, result interface{}, options ...gormq.Option) error
	UpdateById(ctx context.Context, id uint, record *T, options ...gormq.Option) (*T, error)
	Update(ctx context.Context, record *T, options ...gormq.Option) (*T, error)
	Count(ctx context.Context, options ...gormq.Option) (int64, error)
	Relay(ctx context.Context, options ...relay.PaginateOption) (*relay.Connection[T], error)

	// transaction
	Tx(*gorm.DB) IBaseRepository[T]
	UseTxIfExist(ctx context.Context) IBaseRepository[T]
}

type BaseRepository[T any] struct {
	db    *gorm.DB
	model *T
}

func NewBaseRepository[T any](db *gorm.DB) IBaseRepository[T] {
	var model T
	return &BaseRepository[T]{db, &model}
}

func (repo *BaseRepository[T]) Tx(tx *gorm.DB) IBaseRepository[T] {
	copy := &BaseRepository[T]{}
	copy.db = tx
	return copy
}

func applyOptions(q *gorm.DB, options []gormq.Option) *gorm.DB {
	for _, opt := range options {
		q = opt(q)
	}
	return q
}

func (repo *BaseRepository[T]) GetById(ctx context.Context, id uint, options ...gormq.Option) (*T, error) {
	var resource T
	query := repo.db.WithContext(ctx).Model(repo.model).Where("id = ?", id)
	query = applyOptions(query, options)
	if err := query.First(&resource).Error; err != nil {
		return nil, err
	}
	return &resource, nil
}

func (repo *BaseRepository[T]) GetMany(ctx context.Context, options ...gormq.Option) ([]T, error) {
	var resources []T

	query := repo.db.WithContext(ctx).Model(repo.model)
	query = applyOptions(query, options)
	if err := query.Find(&resources).Error; err != nil {
		return nil, err
	}
	return resources, nil
}

func (repo *BaseRepository[T]) GetManyTo(ctx context.Context, result interface{}, options ...gormq.Option) error {
	query := repo.db.WithContext(ctx)
	query = applyOptions(query, options)
	return query.Find(result).Error
}

func (repo *BaseRepository[T]) Create(ctx context.Context, record *T) (*T, error) {
	err := repo.db.WithContext(ctx).Create(record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (repo *BaseRepository[T]) CreateMany(ctx context.Context, records []T) ([]T, error) {
	err := repo.db.WithContext(ctx).Create(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (repo *BaseRepository[T]) UpdateById(ctx context.Context, id uint, record *T, options ...gormq.Option) (*T, error) {
	query := repo.db.WithContext(ctx).Model(record).Where("id = ?", id)
	query = applyOptions(query, options)
	if err := query.Updates(record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (repo *BaseRepository[T]) Update(ctx context.Context, record *T, options ...gormq.Option) (*T, error) {
	query := repo.db.WithContext(ctx).Model(record)
	query = applyOptions(query, options)
	if err := query.Updates(record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (repo *BaseRepository[T]) Delete(ctx context.Context, options ...gormq.Option) error {
	query := repo.db.WithContext(ctx).Model(repo.model)
	query = applyOptions(query, options)
	return query.Delete(repo.model).Error
}

// DeleteById will soft delete the record, If you want to hard delete, use WithHardDelete option
//
// Example:
// DeleteById(ctx, id, &models.User{}, WithHardDelete())
func (repo *BaseRepository[T]) DeleteById(ctx context.Context, id uint, options ...gormq.Option) error {
	query := repo.db.WithContext(ctx).Model(repo.model).Where("id = ?", id)
	query = applyOptions(query, options)
	return query.Delete(id).Error
}

func (repo *BaseRepository[T]) GetOne(ctx context.Context, options ...gormq.Option) (*T, error) {
	var resource T

	query := repo.db.WithContext(ctx).Model(repo.model)
	query = applyOptions(query, options)
	if err := query.First(&resource).Error; err != nil {
		return nil, err
	}

	return &resource, nil
}

func (repo *BaseRepository[T]) GetOneOrCreate(ctx context.Context, record *T, option gormq.Option, options ...gormq.Option) (isNew bool, resource *T, err error) {

	resource, err = repo.GetOne(ctx, gormq.WithLockUpdate(), option, gormq.Multi(options...))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resource, err = repo.Create(ctx, record)
			if err != nil {
				return false, nil, err
			}

			return true, resource, nil
		}
	}

	return
}

func (repo *BaseRepository[T]) GetOneTo(ctx context.Context, result interface{}, options ...gormq.Option) error {
	query := repo.db.WithContext(ctx)
	query = applyOptions(query, options)
	if err := query.First(&result).Error; err != nil {
		return err
	}

	return nil
}

func (repo *BaseRepository[T]) Relay(ctx context.Context, options ...relay.PaginateOption) (*relay.Connection[T], error) {
	defaultOptions := []relay.PaginateOption{
		relay.MaxLimit(100),
	}
	defaultOptions = append(defaultOptions, options...)
	query := repo.db.WithContext(ctx).Model(repo.model)

	return relay.Paginate[T](query, defaultOptions...)
}

func (repo *BaseRepository[T]) Count(ctx context.Context, options ...gormq.Option) (int64, error) {
	var total int64
	query := repo.db.WithContext(ctx).Model(repo.model)
	query = applyOptions(query, options)
	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *BaseRepository[T]) GetManyAndCount(ctx context.Context, resp any, count *int64, options ...gormq.Option) error {
	isGetManyOnly := resp != nil
	isGetCountOnly := count != nil
	isGetManyOnlyAndCount := resp != nil && count != nil

	if isGetManyOnlyAndCount {
		return repo.GetManyTo(ctx, resp, gormq.Multi(options...), gormq.WithCount(count, gormq.WithModel(repo.model)))
	}

	if isGetCountOnly {
		total, err := repo.Count(ctx, gormq.Multi(options...), gormq.WithoutLimitAndOffset())
		if err != nil {
			return err
		}
		*count = total
		return nil
	}

	if isGetManyOnly {
		return repo.GetManyTo(ctx, resp, gormq.WithModel(repo.model), gormq.Multi(options...))
	}

	return errors.New("unsupport operation resp or count must not be nil")
}

func (repo *BaseRepository[T]) UseTxIfExist(ctx context.Context) IBaseRepository[T] {
	instance := dbtx.GetTx(ctx)
	if instance == nil {
		return repo
	}
	return repo.Tx(instance)
}
