package dto

import (
	"backend/core/utils/json"
	"backend/internal/pagi"
)

// Requests & responses Data Transfer Object
type CreateConfigurationRequest struct {
	Key         string    `json:"key" validate:"required"`
	Value       json.JSON `json:"value,omitempty" validate:"excluded_if=Type group"`
	Type        string    `json:"type" validate:"required,oneof=string textarea int float bool array link date datetime dateRange json image multiImage group object"`
	Description string    `json:"description,omitempty"`
	Metadata    json.JSON `json:"metadata,omitempty"`

	// Parent Configuration
	ParentId *uint `json:"parentId" validate:"required_unless=Type group"`
	RootId   *uint `json:"rootId" validate:"required_unless=Type group"`
}

type GetConfigurationRequest struct {
	ID uint `param:"id" validate:"required"`
}

type GetConfigurationByKeyRequest struct {
	Key string `param:"key" validate:"required"`
}

type UpdateConfigurationRequest struct {
	Key         string    `json:"key"`
	Value       json.JSON `json:"value"`
	Type        string    `json:"type" validate:"required,oneof=string textarea int float bool array link date datetime dateRange json image multiImage group object"`
	Description string    `json:"description,omitempty"`
	Metadata    json.JSON `json:"metadata,omitempty"`
	ParentId    uint      `json:"parentId" validate:"required_with=Value"`
	RootId      *uint     `json:"rootId" validate:"required_unless=Type group"`
}

// not nil will update
type UpdatePatchConfigurationRequest struct {
	Key         *string    `json:"key,omitempty"`
	Value       *json.JSON `json:"value,omitempty"`
	Type        *string    `json:"type,omitempty" validate:"required,oneof=string textarea int float bool array link date datetime dateRange json image multiImage group object"`
	Description string     `json:"description,omitempty"`
	Metadata    *json.JSON `json:"metadata,omitempty"`
}

type GetManyQuery struct {
	pagi.QueryDto
}
