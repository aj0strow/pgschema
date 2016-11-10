package psql

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/temp"
	"github.com/davecgh/go-spew/spew"
)

const userTable = `
	CREATE TABLE users (
		name text
	)
`

func TestLoadIndexes(t *testing.T) {
	type Test struct {
		Setup string
		Index db.Index
	}
	tests := []Test{
		Test{
			`CREATE INDEX myindex ON users (name)`,
			db.Index{
				TableName: "users",
				IndexName: "myindex",
			},
		},
		Test{
			`CREATE UNIQUE INDEX users_name_key ON users (name)`,
			db.Index{
				TableName: "users",
				IndexName: "users_name_key",
				Unique:    true,
			},
		},
		Test{
			`ALTER TABLE users ADD PRIMARY KEY (name)`,
			db.Index{
				TableName: "users",
				IndexName: "users_pkey",
				Unique:    true,
				Primary:   true,
				Exprs:     []string{"name"},
			},
		},
	}
	for _, test := range tests {
		runIndexTest(t, test.Setup, test.Index)
	}
}

func runIndexTest(t *testing.T, q string, ix db.Index) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if _, err := conn.Exec(userTable); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Exec(q); err != nil {
		t.Fatal(err)
	}
	tableName := "users"
	indexes, err := LoadIndexes(conn, conn.SchemaName, tableName)
	if err != nil {
		t.Fatal(err)
	}
	if len(indexes) != 1 {
		t.Fatalf("invalid index count: %d", len(indexes))
	}
	index := indexes[0]
	if !reflect.DeepEqual(index, ix) {
		t.Errorf("invalid index")
		t.Errorf(q)
		spew.Dump(ix, index)
	}
}
