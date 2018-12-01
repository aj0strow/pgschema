package psql

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/temp"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestLoadTables(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Exec(`CREATE TABLE users ()`)
	if err != nil {
		t.Fatal(err)
	}
	tables, err := LoadTables(conn, conn.SchemaName)
	if err != nil {
		t.Fatal(err)
	}
	expected := []*db.Table{
		&db.Table{
			TableName: "users",
		},
	}
	if !reflect.DeepEqual(tables, expected) {
		t.Errorf("wrong table list")
		spew.Dump(tables, expected)
	}
}
