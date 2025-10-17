package module_template

import "strings"

var Model = strings.ReplaceAll(`package model

import "backend/app/common/constant/table"

type {{.EntityPascal}} struct {
	BaseModel
	Title string $json:"title"$
}

func (t *{{.EntityPascal}}) TableName() string {
	return table.{{.EntityPascal}}
}

func (t *{{.EntityPascal}}) New{{.EntityPascal}}Model() *{{.EntityPascal}} {
	return &{{.EntityPascal}}{
		BaseModel: NewBaseModel(),
		Title:       t.Title,
	}
}
`,
	"$",
	"`",
)
