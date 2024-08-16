package main

import (

	// #inject:moduleImport (do not remove this comment, it is used by the code generator)

	"github.com/gva/api/admin"
	_ "github.com/gva/internal/ent/runtime"

	_ "github.com/swaggo/swag"
)

func main() {
	admin.Run()
}
