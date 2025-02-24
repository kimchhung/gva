package module_template

import "strings"

var Model = strings.ReplaceAll(`package model

import (
	"backend/app/common/constant/table"
)

type {{.EntityPascal}} struct {
	BaseModelV2
	Title string $json:"title"$
}

func (t *{{.EntityPascal}}) TableName() string {
	return constant.{{.EntityPascal}}TableName
}

func (t *{{.EntityPascal}}) New{{.EntityPascal}}Model() *{{.EntityPascal}} {
	return &{{.EntityPascal}}{
		BaseModelV2: NewBaseModelV2(),
		Title:       t.Title,
	}
}
`,
	"$",
	"`",
)
