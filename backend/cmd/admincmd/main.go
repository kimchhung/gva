package main

import (
	"github.com/kimchhung/gva/backend-echo/cmd/admincmd/cmd"
	_ "github.com/kimchhung/gva/backend-echo/internal/ent/runtime"
)

func main() {
	// go run main push | pull
	cmd.Execute()
}
