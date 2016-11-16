package cli

import (
	"flag"
	"github.com/aj0strow/pgschema/order"
	"github.com/aj0strow/pgschema/source/hcl"
	"github.com/aj0strow/pgschema/source/psql"
	"github.com/jackc/pgx"
	"io/ioutil"
	"strings"
)

type Update struct {
	Term
}

func (cmd *Update) Synopsis() string {
	return `Update database schema to match source file`
}

func (cmd *Update) Run(args []string) int {
	var (
		source string
		uri    string
		dsn    string
	)
	f := flag.NewFlagSet("update", flag.ContinueOnError)
	f.StringVar(&source, "source", "", "")
	f.StringVar(&uri, "uri", "", "")
	f.StringVar(&dsn, "dsn", "", "")
	if err := f.Parse(args); err != nil {
		return 1
	}
	if source == "" {
		cmd.Error("-source input file can't be empty")
		return 1
	}
	data, err := ioutil.ReadFile(source)
	if err != nil {
		cmd.Error(err.Error())
		return 1
	}
	a, err := hcl.ParseBytes(data)
	if err != nil {
		cmd.Error(err.Error())
		return 1
	}
	var (
		pgConfig *pgx.ConnConfig
	)
	if dsn != "" {
		config, err := pgx.ParseDSN(dsn)
		if err != nil {
			cmd.Error(err.Error())
			return 1
		}
		pgConfig = &config
	}
	if uri != "" {
		config, err := pgx.ParseURI(uri)
		if err != nil {
			cmd.Error(err.Error())
			return 1
		}
		pgConfig = &config
	}
	if pgConfig == nil {
		cmd.Error("database connection required, provide -dsn or -uri")
		return 1
	}
	conn, err := pgx.Connect(*pgConfig)
	if err != nil {
		cmd.Error(err.Error())
		return 1
	}
	defer conn.Close()
	b, err := psql.LoadDatabaseNode(conn)
	if err != nil {
		cmd.Error(err.Error())
		return 1
	}
	tx, err := conn.Begin()
	if err != nil {
		cmd.Error(err.Error())
		return 1
	}
	defer tx.Rollback()

	changes := order.Changes(a, b)
	for _, change := range changes {
		cmd.Info(change.String())
		_, err := tx.Exec(change.String())
		if err != nil {
			cmd.Error(err.Error())
			return 1
		}
	}
	err = tx.Commit()
	if err != nil {
		cmd.Error(err.Error())
		return 1
	}
	return 0
}

func (cmd *Update) Help() string {
	help := `
Usage: pgschema update [options]
 Update database schema to match source file. 
 
Options:
 
  -source=<path>          The source file in hcl schema format.
  
  -uri=<uri>              Postgres database connection uri, eg. 'postgres://'.
  
  -dsn=<dsn>              Postgres database connection dsn, eg. 'dbname='.
`
	return strings.TrimSpace(help)
}
