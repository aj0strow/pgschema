package psql

import (
	"testing"

	"github.com/aj0strow/pgschema/temp"
)

func TestLoadSchemas(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	schemas, err := LoadSchemas(conn)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, schema := range schemas {
		if schema.SchemaName == conn.SchemaName {
			found = true
		}
	}
	if !found {
		t.Fatal("missing temp schema name")
	}
}
