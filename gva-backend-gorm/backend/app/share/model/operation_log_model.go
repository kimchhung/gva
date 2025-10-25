package model

import (
	"backend/app/share/constant/table"
	"time"
)

type OperationLogData map[string]any

type OperationLog struct {
	ID        uint             `json:"id" gorm:"primarykey"`
	CreatedAt time.Time        `json:"createdAt" gorm:"autoCreateTime; type:datetime(3); default:NOW(3)"`
	AdminId   uint             `json:"adminId"`
	Admin     *Admin           `json:"admin" gorm:"foreignKey:AdminId"`
	RoleIds   []uint           `json:"roleIds" gorm:"serializer:json"`
	Method    string           `json:"method"`
	Path      string           `json:"path"`
	Scope     string           `json:"scope"`
	IP        string           `json:"ip"`
	Data      OperationLogData `json:"data" gorm:"serializer:json"`
	Code      int              `json:"code"`
	Error     string           `json:"error"`
	Msg       string           `json:"msg"`
	Latency   int64            `json:"latency"`
}

func (t *OperationLog) TableName() string {
	return table.OperationLog
}
