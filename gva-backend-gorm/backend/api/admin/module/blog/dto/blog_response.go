package dto

import (
	"backend/app/common/constant/table"
	"backend/app/common/model"
)

// Requests & responses Data Transfer Object
type BlogResponse struct {
	model.BaseModel

	Title string `json:"title"`
}

func (BlogResponse) TableName() string {
	return table.Blog
}
