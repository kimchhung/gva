package types

import "backend/app/share/constant/table"

type BaseAdmin struct {
	ID          uint     `json:"id"`
	Status      int      `json:"status"`
	Name        string   `json:"name"`
	Username    string   `json:"username"`
	IpWhiteList []string `json:"ipWhiteList" gorm:"serializer:json"`
}

func (BaseAdmin) TableName() string {
	return table.Admin
}
