package main

import (
	"github.com/mitchellh/cli"
	"os"
)

const (
	Success       = 0
	BadInput      = 1
	DatabaseError = 2
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
	c := cli.NewCLI("pgschema", "1.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"update": func() (cli.Command, error) {
			return &Update{ui}, nil
		},
	}
	exitCode, err := c.Run()
	if err != nil {
		ui.Error(err.Error())
	}
	os.Exit(exitCode)
}
