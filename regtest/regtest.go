package regtest

import (
	"bytes"
	"github.com/aj0strow/pgschema/order"
	"github.com/aj0strow/pgschema/source/hcl"
	"github.com/aj0strow/pgschema/source/psql"
	"github.com/aj0strow/pgschema/temp"
	"io/ioutil"
)

type RegressionTest struct {
	Conn       *temp.Conn
	SetupFile  string
	SourceFile string
}

func (r *RegressionTest) Run() ([]order.Change, error) {
	f1, err := ioutil.ReadFile(r.SetupFile)
	if err != nil {
		return nil, err
	}
	if _, err := r.Conn.Exec(string(f1)); err != nil {
		return nil, err
	}
	f2, err := ioutil.ReadFile(r.SourceFile)
	if err != nil {
		return nil, err
	}
	schema := bytes.Replace(f2, []byte("_schema_"), []byte(r.Conn.SchemaName), 1)
	a, err := hcl.ParseBytes(schema)
	if err != nil {
		return nil, err
	}
	if err := a.Err(); err != nil {
		return nil, err
	}
	b, err := psql.LoadDatabaseNode(r.Conn)
	if err != nil {
		return nil, err
	}
	if err := b.Err(); err != nil {
		return nil, err
	}
	return order.Changes(a, b), nil
}
