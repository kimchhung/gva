package model

import (
	"backend/app/share/constant/table"
	"time"
)

type Permission struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Group     string    `json:"group,omitempty"`
	Name      string    `json:"name,omitempty"`
	Scope     string    `json:"scope,omitempty"`
	Order     int       `json:"order,omitempty"`
}

func (Permission) TableName() string {
	return table.Permission
}
