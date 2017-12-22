package main

import (
	"flag"
	"github.com/aj0strow/pgschema/order"
	"github.com/aj0strow/pgschema/source/hcl"
	"github.com/aj0strow/pgschema/source/psql"
	"github.com/jackc/pgx"
	"github.com/mitchellh/cli"
	"io/ioutil"
)

type Update struct {
	cli.Ui
}

func (cmd *Update) Synopsis() string {
	return "Update database schema to match source file."
}

func (cmd *Update) Help() string {
	return `
 => Update database schema to match source file. 

pgschema update [flags]
 
  -source <path>          Path to source file in HCL schema format.
  
  -uri <uri>              Postgres database connection uri, eg. 'postgres://'.
  
  -dsn <dsn>              Postgres database connection dsn, eg. 'dbname='.
  
  -dryRun                 Print SQL update commands without executing.
`
}

func (cmd *Update) Run(args []string) int {
	var (
		source string
		uri    string
		dsn    string
		dryRun bool
	)
	f := flag.NewFlagSet("update", flag.ContinueOnError)
	f.StringVar(&source, "source", "", "")
	f.StringVar(&uri, "uri", "", "")
	f.StringVar(&dsn, "dsn", "", "")
	f.BoolVar(&dryRun, "dryRun", false, "")
	if err := f.Parse(args); err != nil {
		cmd.Error(err.Error())
		return BadInput
	}
	if source == "" {
		cmd.Error("-source input file can't be empty")
		return BadInput
	}
	data, err := ioutil.ReadFile(source)
	if err != nil {
		cmd.Error(err.Error())
		return BadInput
	}
	a, err := hcl.ParseBytes(data)
	if err != nil {
		cmd.Error(err.Error())
		return BadInput
	}
	if err := a.Err(); err != nil {
		cmd.Error(err.Error())
		return BadInput
	}
	var pgConfig *pgx.ConnConfig
	if dsn != "" {
		config, err := pgx.ParseDSN(dsn)
		if err != nil {
			cmd.Error(err.Error())
			return BadInput
		}
		pgConfig = &config
	}
	if uri != "" {
		config, err := pgx.ParseURI(uri)
		if err != nil {
			cmd.Error(err.Error())
			return BadInput
		}
		pgConfig = &config
	}
	if pgConfig == nil {
		cmd.Error("database connection required, provide -dsn or -uri")
		return BadInput
	}
	conn, err := pgx.Connect(*pgConfig)
	if err != nil {
		cmd.Error(err.Error())
		return DatabaseError
	}
	defer conn.Close()
	b, err := psql.LoadDatabaseNode(conn)
	if err != nil {
		cmd.Error(err.Error())
		return DatabaseError
	}
	changes := order.Changes(a, b)
	if dryRun {
		for _, change := range changes {
			cmd.Info(change.String())
		}
		return Success
	}
	tx, err := conn.Begin()
	if err != nil {
		cmd.Error(err.Error())
		return DatabaseError
	}
	defer tx.Rollback()
	for _, change := range changes {
		cmd.Info(change.String())
		_, err := tx.Exec(change.String())
		if err != nil {
			cmd.Error(err.Error())
			return DatabaseError
		}
	}
	if err := tx.Commit(); err != nil {
		cmd.Error(err.Error())
		return DatabaseError
	}
	return Success
}
