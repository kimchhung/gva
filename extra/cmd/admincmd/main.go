package main

import (
	"github.com/kimchhung/gva/extra/cmd/admincmd/cmd"
	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
)

func main() {
	// go run main push | pull
	cmd.Execute()
}
