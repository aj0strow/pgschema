package t001

import (
	"bytes"
	"github.com/aj0strow/pgschema/order"
	"github.com/aj0strow/pgschema/source/hcl"
	"github.com/aj0strow/pgschema/source/psql"
	"github.com/aj0strow/pgschema/temp"
	"github.com/kr/pretty"
	"io/ioutil"
	"testing"
)

func TestNumericDouble(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	file, err := ioutil.ReadFile("./db.sql")
	if err != nil {
		t.Fatal(err)
	}
	query := bytes.Replace(file, []byte("_schema_"), []byte(conn.SchemaName), 1)
	if _, err := conn.Exec(string(query)); err != nil {
		t.Fatal(err)
	}

	file2, err := ioutil.ReadFile("./schema.hcl")
	if err != nil {
		t.Fatal(err)
	}
	schema := bytes.Replace(file2, []byte("_schema_"), []byte(conn.SchemaName), 1)

	a, err := hcl.ParseBytes(schema)
	if err != nil {
		t.Fatal(err)
	}
	if err := a.Err(); err != nil {
		t.Fatal(err)
	}

	b, err := psql.LoadDatabaseNode(conn)
	if err != nil {
		t.Fatal(err)
	}
	if err := b.Err(); err != nil {
		t.Fatal(err)
	}

	pretty.Log(a, b)

	have := order.Changes(a, b)

	want := []order.Change{
		order.AlterTable{
			SchemaName: conn.SchemaName,
			TableName:  "accounts",
			Change: order.AlterColumn{
				ColumnName: "balance",
				Change: order.SetDataType{
					DataType: "double precision",
				},
			},
		},
	}

	for _, wc := range want {
		found := false
		for _, hc := range have {
			if wc.String() == hc.String() {
				found = true
			}
		}
		if !found {
			t.Errorf(wc.String())
		}
	}
}
