package module_template

var DtoRequest = `package dto

import "github.com/kimchhung/gva/backend/internal/ent"

// Requests & responses Data Transfer Object
type {{.EntityPascal}}Request struct {
	*ent.{{.EntityPascal}}
}
`

var DtoResponse = `package dto

import "github.com/kimchhung/gva/backend/internal/ent"

// Requests & responses Data Transfer Object
type {{.EntityPascal}}Response struct {
	*ent.{{.EntityPascal}}
}
`
