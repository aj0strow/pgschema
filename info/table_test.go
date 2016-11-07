package info

import (
	"github.com/aj0strow/pgschema/temp"
	"reflect"
	"testing"
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
	expected := []Table{
		Table{
			TableName: "users",
		},
	}
	if !reflect.DeepEqual(tables, expected) {
		t.Errorf("wrong table list: %+v", tables)
	}
}