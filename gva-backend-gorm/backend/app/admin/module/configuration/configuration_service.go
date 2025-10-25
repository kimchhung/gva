package configuration

import (
	"backend/app/admin/module/configuration/dto"
	"backend/app/share/constant/cache"
	apperror "backend/app/share/error"
	"backend/app/share/model"
	repository "backend/app/share/repository"
	"backend/app/share/service"
	"backend/core/database"
	"backend/core/utils"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ConfigurationService struct {
	repo    *repository.ConfigurationRepo
	redis_s *service.RedisService
}

// NewConfigurationService initializes a new ConfigurationService with a JwtService and a UserStore.
func NewConfigurationService(repo *repository.ConfigurationRepo, redis_s *service.RedisService) *ConfigurationService {
	return &ConfigurationService{
		repo:    repo,
		redis_s: redis_s,
	}
}

// CreateConfiguration creates a new Configuration.
func (s *ConfigurationService) CreateConfiguration(ctx context.Context, p *dto.CreateConfigurationRequest) (*dto.ConfigurationResponse, error) {
	body := utils.MustCopy(new(model.Configuration), p)
	// Default base model
	body.BaseModel = model.NewBaseModel()
	created, err := s.repo.Create(ctx, body)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, apperror.ErrDuplicatedRecord
		}

		return nil, err
	}

	// Save to cache
	cacheKey := getRedisKeyConfig(created.Key)
	err = s.redis_s.SetJson(ctx, cacheKey, created, cache.AGE_ONE_MONTH)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.ConfigurationResponse), created), nil
}

// GetConfiguration gets a Configuration by ID.
func (s *ConfigurationService) GetConfiguration(ctx context.Context, id uint) (*dto.ConfigurationResponse, error) {
	configuration, err := s.repo.GetById(ctx, id,
		gormq.Preload("Children"),
		gormq.Preload("Parent"),
	)
	if err != nil {
		// Check if the error is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}

		return nil, err
	}

	// Save to cache
	cacheKey := getRedisKeyConfig(configuration.Key)
	err = s.redis_s.SetJson(ctx, cacheKey, configuration, cache.AGE_ONE_MONTH)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.ConfigurationResponse), configuration), nil
}

// LockForUpdate locks a Configuration for update.
func (s *ConfigurationService) LockForUpdate(ctx context.Context, id uint) database.TxOperaton {
	return func(tx *gorm.DB) error {
		_, err := s.repo.Tx(tx).GetById(ctx, id, gormq.WithSelect("id", "description", "value"), gormq.WithLockUpdate())
		if err != nil {
			// Check if the error is a not found error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				panic(apperror.ErrNotFound)
			}

			return err
		}

		return nil
	}
}

// UpdateConfiguration updates a Configuration.
func (s *ConfigurationService) UpdateConfiguration(ctx context.Context, id uint, p *dto.UpdateConfigurationRequest) (updatedRes *dto.ConfigurationResponse, err error) {
	err = s.repo.MultiTransaction(
		s.LockForUpdate(ctx, id),
		func(tx *gorm.DB) error {
			body := utils.MustCopy(new(model.Configuration), p)
			body.ID = id

			updated, err := s.repo.Tx(tx).Update(ctx, body)
			if err != nil {
				return err
			}

			updatedRes = utils.MustCopy(new(dto.ConfigurationResponse), updated)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return s.GetConfiguration(ctx, id)
}

// UpdateConfiguration updates a Configuration.
func (s *ConfigurationService) UpdatePatchConfiguration(ctx context.Context, id uint, p *dto.UpdatePatchConfigurationRequest) (resp map[string]any, err error) {
	err = s.repo.MultiTransaction(
		s.LockForUpdate(ctx, id),
		func(tx *gorm.DB) error {
			columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{})
			dbCols, res := utils.StructToMap(p, columnMap)
			resp = res

			if len(dbCols) > 0 {
				return tx.Model(&model.Configuration{}).
					Scopes(
						gormq.Where(gormq.Equal("id", id)),
					).
					Updates(dbCols).Error
			}

			return apperror.NewError(
				apperror.ErrBadRequest,
				apperror.Join(
					fmt.Errorf("required at least one field to update, support fields: %s", columnMap.Keys()),
				))
		},
	)
	return
}

// DeleteConfiguration deletes a Configuration by ID.
func (s *ConfigurationService) DeleteConfiguration(ctx context.Context, id uint) error {
	// Check if the Configuration exists
	config, err := s.GetConfiguration(ctx, id)
	if err != nil {
		return err
	}

	// Set all the children to parent_id = NULL
	_, err = s.repo.Update(ctx, &model.Configuration{ParentId: nil}, gormq.Where(gormq.Equal("parent_id", config.ID)))
	if err != nil {
		return err
	}

	// Delete the Configuration
	err = s.repo.DeleteById(ctx, id)
	if err != nil {
		return err
	}

	// Delete from cache
	cacheKey := getRedisKeyConfig(config.Key)
	_, err = s.redis_s.Del(ctx, cacheKey)
	if err != nil {
		return err
	}

	return nil
}

// GetConfigurations gets all Configurations.
func (s *ConfigurationService) GetConfigurations(ctx context.Context, query *dto.GetManyQuery) ([]dto.ConfigurationResponse, *pagi.MetaDto, error) {
	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"id":        gormq.Ignore(),
		"key":       gormq.Ignore(),
		"type":      gormq.Ignore(),
		"createdAt": gormq.ToSnake(),
	})

	searchColumns := columnMap.Pick(
		"id",
		"key",
		"type",
	).Values()

	resp, respMeta := pagi.PrepareResponse[dto.ConfigurationResponse](&query.QueryDto)
	err := s.repo.GetManyAndCount(ctx, &resp, respMeta.TotalCount,
		gormq.WithPageAndLimit(query.Page, query.Limit),
		gormq.Where(gormq.WithFilters(query.Filters, columnMap)),
		gormq.WithSorts(query.Sorts, columnMap),
		gormq.WithSearch(query.Search, searchColumns...),
	)

	if err != nil {
		return nil, nil, err
	}
	return resp, respMeta, nil
}

// GetConfigurationByKey gets a Configuration by key.
// TODO: Implement Caching
func (s *ConfigurationService) GetConfigurationByKey(ctx context.Context, key string) (*dto.ConfigurationResponse, error) {
	configuration, err := s.repo.UseTxIfExist(ctx).GetOne(ctx,
		gormq.Where(
			gormq.Equal("configurations.key", key),
		),
		gormq.Preload("Children"),
		gormq.Preload("Parent"),
	)
	if err != nil {
		// Check if the error is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}

		return nil, err
	}

	return utils.MustCopy(new(dto.ConfigurationResponse), configuration), nil
}

func (s *ConfigurationService) GetDocsConfiguration(ctx context.Context) ([]dto.ConfigurationResponse, error) {
	var resp []dto.ConfigurationResponse
	err := s.repo.GetManyTo(ctx, &resp,
		gormq.WithModel(&model.Configuration{}),
		gormq.Where(
			gormq.Equal("configurations.key", "document"),
		),
		gormq.Preload("AllChildren"),
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getRedisKeyConfig(key string) string {
	return cache.KeyConfig.WithKeyValue("key", key).String()
}
