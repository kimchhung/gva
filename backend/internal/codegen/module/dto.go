package module_template

var DtoRequest = `package dto

// Requests Data Transfer Object
type {{.EntityPascal}}Request struct {
}
`

var DtoResponse = `package dto

import "github.com/gva/internal/ent"

// Responses Data Transfer Object
type {{.EntityPascal}}Response struct {
}
`
