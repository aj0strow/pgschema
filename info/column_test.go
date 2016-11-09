package info

import (
	"fmt"
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/temp"
	"reflect"
	"testing"
)

func TestLoadColumn(t *testing.T) {
	type Test struct {
		Query  string
		Column db.Column
	}
	tests := []Test{
		Test{
			`name text`,
			db.Column{
				ColumnName: "name",
				DataType:   "text",
			},
		},
	}
	for _, test := range tests {
		runLoadColumn(t, test.Query, test.Column)
	}
}

func runLoadColumn(t *testing.T, q string, c db.Column) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tableName := "test"
	_, err = conn.Exec(fmt.Sprintf(`CREATE TABLE %s (%s)`, tableName, q))
	if err != nil {
		t.Fatal(err)
	}
	cs, err := LoadColumns(conn, conn.SchemaName, tableName)
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
