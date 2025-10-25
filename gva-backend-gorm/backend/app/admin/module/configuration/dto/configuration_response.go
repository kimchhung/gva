package dto

import (
	"backend/app/share/constant/table"
	"backend/app/share/model"
	"backend/core/utils/json"
)

// Requests & responses Data Transfer Object
type ConfigurationResponse struct {
	model.BaseModel

	Key         string    `json:"key"`
	Value       json.JSON `json:"value,omitempty" gorm:"type:json"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	Metadata    json.JSON `json:"metadata,omitempty" gorm:"type:json"`

	// Parent Configuration
	ParentId *uint                  `json:"parentId,omitempty"`
	Parent   *ConfigurationResponse `json:"parent,omitempty" gorm:"foreignKey:ParentId;references:ID"`

	RootId *uint                  `json:"rootId,omitempty"`
	Root   *ConfigurationResponse `json:"root,omitempty" gorm:"foreignKey:RootId;references:ID"`

	// Children Configuration
	Children    []ConfigurationResponse `json:"children,omitempty" gorm:"foreignKey:ParentId;references:ID"`
	AllChildren []ConfigurationResponse `json:"allChildren,omitempty" gorm:"foreignKey:RootId;references:ID"`
}

func (ConfigurationResponse) TableName() string {
	return table.Configuration
}
