package main

import (
	"github.com/kimchhung/gva/backend/cmd/admincmd/cmd"
	_ "github.com/kimchhung/gva/backend/internal/ent/runtime"
)

func main() {
	// go run main push | pull
	cmd.Execute()
}
