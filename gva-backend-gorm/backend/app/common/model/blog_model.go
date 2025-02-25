package model

import "backend/app/common/constant/table"

type Blog struct {
	BaseModel
	Title string `json:"title"`
}

func (t *Blog) TableName() string {
	return table.Blog
}

func (t *Blog) NewBlogModel() *Blog {
	return &Blog{
		BaseModel: NewBaseModel(),
		Title:       t.Title,
	}
}
