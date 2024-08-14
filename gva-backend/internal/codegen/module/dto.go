package module_template

import "strings"

var DtoRequest = strings.ReplaceAll(`package dto
import "github.com/gva/internal/rql"

// Requests Data Transfer Object
type {{.EntityPascal}}Request struct {
	Name string $json:"name"$
}

type {{.EntityPascal}}PagedRequest struct {
	rql.Params
	Selects string $query:"selects"$
}
`,
	"$",
	"`",
)

var DtoResponse = strings.ReplaceAll(`package dto
import 	"github.com/gva/internal/ent"

// Responses Data Transfer Object
type {{.EntityPascal}}Response struct {
	*ent.{{.EntityPascal}}
}
`, "$", "`")
