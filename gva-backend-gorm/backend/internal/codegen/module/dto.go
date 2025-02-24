package module_template

import "strings"

var DtoRequest = strings.ReplaceAll(`package dto

import (
	"backend/internal/pagi"
)

// Requests & responses Data Transfer Object
type Create{{.EntityPascal}}Request struct {
	Title string $json:"title" validate:"required"$
}

type Get{{.EntityPascal}}Request struct {
	ID uint $param:"id" validate:"required"$
}

type Update{{.EntityPascal}}Request struct {
	Title string $json:"title"$
}


// not nil will update
type UpdatePatch{{.EntityPascal}}Request struct {
	Title *string $json:"title"$
}

type GetManyQuery struct {
	pagi.QueryDto
}
`,
	"$",
	"`",
)

var DtoResponse = strings.ReplaceAll(`package dto

import (
	"backend/app/common/constant/table"
	"backend/app/common/model"
)

// Requests & responses Data Transfer Object
type {{.EntityPascal}}Response struct {
	model.BaseModelV2

	Title string $json:"title"$
}

func ({{.EntityPascal}}Response) TableName() string {
	return table.{{.EntityPascal}}
}
`, "$", "`")
