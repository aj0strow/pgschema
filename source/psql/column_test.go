package psql

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/temp"
	"github.com/davecgh/go-spew/spew"
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
		Test{
			`name text NOT NULL`,
			db.Column{
				ColumnName: "name",
				DataType:   "text",
				NotNull:    true,
			},
		},
		Test{
			`name text DEFAULT 'nobody'`,
			db.Column{
				ColumnName: "name",
				DataType:   "text",
				Default:    `'nobody'::text`,
			},
		},
		Test{
			`balance numeric(11,2) DEFAULT 0`,
			db.Column{
				ColumnName:       "balance",
				DataType:         "numeric",
				Default:          "0",
				NumericPrecision: 11,
				NumericScale:     2,
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
		t.Errorf("invalid column")
		t.Errorf(q)
		spew.Dump(c, column)
	}
}
