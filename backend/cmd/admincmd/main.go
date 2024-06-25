package main

import (
	"github.com/gva/cmd/admincmd/cmd"
	_ "github.com/gva/internal/ent/runtime"
)

func main() {
	// go run main push | pull
	cmd.Execute()
}
