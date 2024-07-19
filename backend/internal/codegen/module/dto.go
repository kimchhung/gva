package module_template

var DtoRequest = `package dto

// Requests Data Transfer Object
type {{.EntityPascal}}Request struct {
	Name string
}
`

var DtoResponse = `package dto

// Responses Data Transfer Object
type {{.EntityPascal}}Response struct {
	Name string
}
`
