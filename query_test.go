package pgschema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadTables(t *testing.T) {
	conn, err := Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Exec(`CREATE TABLE users ()`)
	if err != nil {
		t.Fatal(err)
	}
	tables, err := LoadTables(conn, conn.Schema)
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

func TestLoadColumn(t *testing.T) {
	type Test struct {
		Query  string
		Column Column
	}
	tests := []Test{
		Test{
			`name text`,
			Column{
				ColumnName: "name",
				DataType:   "text",
			},
		},
	}
	for _, test := range tests {
		runLoadColumn(t, test.Query, test.Column)
	}
}

func runLoadColumn(t *testing.T, q string, c Column) {
	conn, err := Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	schemaName := conn.Schema
	tableName := "test"
	_, err = conn.Exec(fmt.Sprintf(`CREATE TABLE %s (%s)`, tableName, q))
	if err != nil {
		t.Fatal(err)
	}
	cs, err := LoadColumns(conn, schemaName, tableName)
	if err != nil {
		t.Fatal(err)
	}
	if len(cs) != 1 {
		t.Fatalf("invalid column count: %d", len(cs))
	}
	column := cs[0]
	if !reflect.DeepEqual(column, c) {
		fmt.Printf("want %+v\n", c)
		fmt.Printf("have %+v\n", column)
		t.Errorf("invalid column for %s", q)
	}
}
