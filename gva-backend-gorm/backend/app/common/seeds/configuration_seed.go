package seeds

import (
	"backend/app/common/model"
	"backend/internal/bootstrap/database"
	"backend/utils/json"
	"context"

	"gorm.io/gorm"
)

type ConfigurationSeeder struct {
}

func NewConfigurationSeeder() database.Seeder {
	return &ConfigurationSeeder{}
}

func (s ConfigurationSeeder) Name() string {
	return "ConfigurationSeeder"
}

func (s ConfigurationSeeder) Count(ctx context.Context, db *gorm.DB) (int, error) {
	var total int64
	if err := db.Model(model.Configuration{}).Where(model.Configuration{
		Key: "document",
	}).Count(&total).Error; err != nil {
		return 0, err
	}

	return int(total), nil
}

func (s ConfigurationSeeder) Seed(ctx context.Context, db *gorm.DB) error {
	document := &model.Configuration{
		BaseModel:   model.NewBaseModel(),
		Key:         "document",
		Description: "Document",
		Type:        model.ValueTypeString,
		Metadata:    json.MustJSON(map[string]any{"icon": "lucide:book-text", "labelEn": "Document", "labelZh": "文件"}),
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(document).Error; err != nil {
			return err
		}

		backendStructure := model.Configuration{
			BaseModel:   model.NewBaseModel(),
			Key:         "document_backend_structure",
			Description: "backend folder structure",
			Type:        model.ValueTypeString,
			Metadata:    json.MustJSON(map[string]any{"icon": "lucide:folder", "labelEn": "Structure", "labelZh": "后端结构"}),
			Value:       json.MustJSON("http://localhost:4000/admin/v1/file/docs/backend/directory.md"),
			RootId:      &document.ID,
		}

		backendNameConvention := model.Configuration{
			BaseModel:   model.NewBaseModel(),
			Key:         "document_backend_name_convention",
			Description: "backend name convention guide",
			Type:        model.ValueTypeString,
			Metadata:    json.MustJSON(map[string]any{"icon": "lucide:book-text", "labelEn": "Name Convention", "labelZh": "名称约定"}),
			Value:       json.MustJSON("http://localhost:4000/admin/v1/file/docs/backend/name-convention.md"),
			RootId:      &document.ID,
		}

		backendCodegen := model.Configuration{
			BaseModel:   model.NewBaseModel(),
			Key:         "document_code_gen",
			Description: "backend folder structure",
			Type:        model.ValueTypeString,
			Metadata:    json.MustJSON(map[string]any{"icon": "lucide:book-text", "labelEn": "Codegen", "labelZh": "Codegen"}),
			Value:       json.MustJSON("http://localhost:4000/admin/v1/file/docs/backend/codegen.md"),
			RootId:      &document.ID,
		}

		backendMigration := model.Configuration{
			BaseModel:   model.NewBaseModel(),
			Key:         "document_backend_migration",
			Description: "backend migration guide",
			Type:        model.ValueTypeString,
			Metadata:    json.MustJSON(map[string]any{"icon": "lucide:folder", "labelEn": "Migration", "labelZh": "Migration"}),
			Value:       json.MustJSON("http://localhost:4000/admin/v1/file/docs/backend/migration.md"),
			RootId:      &document.ID,
		}

		backend := &model.Configuration{BaseModel: model.NewBaseModel(),
			Key:         "document_backend",
			Description: "backend group document",
			Type:        model.ValueTypeGroup,
			Metadata:    json.MustJSON(map[string]any{"labelEn": "Backend", "labelZh": "Backend"}),
			Children: []model.Configuration{
				backendStructure, backendNameConvention, backendCodegen, backendMigration,
			},
			RootId:   &document.ID,
			ParentId: &document.ID,
		}

		api := &model.Configuration{BaseModel: model.NewBaseModel(),
			Key:         "document_api",
			Description: "api group document",
			Type:        model.ValueTypeGroup,
			Metadata:    json.MustJSON(map[string]any{"labelEn": "API", "labelZh": "API"}),
			Children: []model.Configuration{
				{BaseModel: model.NewBaseModel(),
					Key:         "document_api_admin",
					Type:        model.ValueTypeString,
					Description: "admin api document",
					Metadata:    json.MustJSON(map[string]any{"labelEn": "API Admin", "labelZh": "API Admin"}),
					Value:       json.MustJSON("http://localhost:4000/admin/v1/docs"),
					RootId:      &document.ID,
				},
				{BaseModel: model.NewBaseModel(),
					Key:         "document_api_bot",
					Type:        model.ValueTypeString,
					Description: "bot api document",
					Metadata:    json.MustJSON(map[string]any{"labelEn": "API Bot", "labelZh": "API Bot"}),
					Value:       json.MustJSON("http://localhost:4000/bot/v1/docs"),
					RootId:      &document.ID,
				},
			},
			RootId:   &document.ID,
			ParentId: &document.ID,
		}

		if err := tx.Create(backend).Error; err != nil {
			return err
		}

		if err := tx.Create(api).Error; err != nil {
			return err
		}

		return nil
	})
}
