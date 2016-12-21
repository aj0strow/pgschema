package main

import (
	"github.com/aj0strow/pgschema/cli"
	"os"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
