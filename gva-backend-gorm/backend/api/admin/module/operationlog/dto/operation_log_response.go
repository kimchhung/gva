package dto

import (
	"backend/app/common/constant/table"
	"backend/app/common/types"
	"time"
)

// Requests & responses Data Transfer Object
type OperationLogResponse struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time       `json:"createdAt" gorm:"autoCreateTime; type:datetime(3); default:NOW(3)"`
	AdminId   uint            `json:"adminId"`
	Admin     types.BaseAdmin `json:"admin" gorm:"foreignKey:AdminId"`
	RoleIds   []uint          `json:"roleIds" gorm:"serializer:json"`
	Method    string          `json:"method"`
	Path      string          `json:"path"`
	Scope     string          `json:"scope"`
	IP        string          `json:"ip"`
	Data      map[string]any  `json:"data" gorm:"serializer:json"`
	Code      int             `json:"code"`
	Error     string          `json:"error"`
	Msg       string          `json:"msg"`
	Latency   int64           `json:"latency"`
}

func (OperationLogResponse) TableName() string {
	return table.OperationLog
}
