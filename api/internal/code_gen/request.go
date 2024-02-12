package code_gen

var request_template = `package request

// Requests & responses for ArticleController & ArticleService
type {{.Entity}}Request struct {
}
`
