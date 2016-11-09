package cli

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/mitchellh/cli"
	"os"
)

type Term interface {
	cli.Ui
}

type SimpleTx struct {
	*pgx.Tx
}

func (tx *SimpleTx) Exec(query string) error {
	_, err := tx.Tx.Exec(query)
	return err
}

func Run(args []string) int {
	c := cli.NewCLI("pgschema", "0.0")
	c.Args = args
	c.Commands = Commands()
	exitCode, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	return exitCode
}

func Commands() map[string]cli.CommandFactory {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
	return map[string]cli.CommandFactory{
		"update": func() (cli.Command, error) {
			update := &Update{
				Term: ui,
			}
			return update, nil
		},
	}
}
