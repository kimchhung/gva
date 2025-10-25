package model

import (
	"backend/app/share/constant/table"
	"backend/core/utils/json"
)

const (
	ValueTypeString   ValueType = "string"
	ValueTypeInt      ValueType = "int"
	ValueTypeFloat    ValueType = "float"
	ValueTypeBool     ValueType = "bool"
	ValueTypeLink     ValueType = "link"
	ValueTypeDate     ValueType = "date"
	ValueTypeDateTime ValueType = "datetime"
	ValueTypeJSON     ValueType = "json"
	ValueTypeGroup    ValueType = "group"
	ValueTypeArray    ValueType = "array"
	ValueTypeObject   ValueType = "object"
)

type ValueType string

type Configuration struct {
	BaseModel
	Key         string    `json:"key" gorm:"unique;not null"`
	Value       json.JSON `json:"value,omitempty" gorm:"type:json"`
	Description string    `json:"description,omitempty"`
	Type        ValueType `json:"type"`
	Metadata    json.JSON `json:"metadata,omitempty" gorm:"type:json"`

	// Parent Configuration
	ParentId *uint          `json:"parentId"`
	Parent   *Configuration `json:"parent" gorm:"foreignKey:ParentId;references:ID"`

	RootId *uint          `json:"rootId"`
	Root   *Configuration `json:"root" gorm:"foreignKey:RootId;references:ID"`

	Children    []Configuration `json:"children,omitempty" gorm:"foreignKey:ParentId;references:ID"`
	AllChildren []Configuration `json:"allChildren,omitempty" gorm:"foreignKey:RootId;references:ID"`
}

func (Configuration) TableName() string {
	return table.Configuration
}
