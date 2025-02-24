package model

import (
	"backend/app/common/model/softDelete"
	"time"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime; type:datetime(3); default:NOW(3)"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime; type:datetime(3); default:NOW(3) ON UPDATE"`

	DeletedAt softDelete.DeletedAt `json:"-" gorm:"softDelete:milli"`
}

func NewBaseModel() BaseModel {
	baseHardDeleteModel := BaseModel{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return baseHardDeleteModel
}
