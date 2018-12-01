package hcl

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestParseBytes(t *testing.T) {
	type Test struct {
		Input    string
		Database *db.Database
	}
	var (
		tests    []Test
		input    string
		database *db.Database
		tables   []*db.Table
	)

	// Test table with column.
	input = `
schema "public" {
	table "users" {
		column "email" {
			type = "text"
		}
	}
}
	`
	tables = []*db.Table{
		&db.Table{
			TableName: "users",
			Columns: []*db.Column{
				&db.Column{
					ColumnName: "email",
					DataType:   "text",
				},
			},
		},
	}
	database = &db.Database{
		Schemas: []*db.Schema{
			&db.Schema{
				SchemaName: "public",
				Tables:     tables,
			},
		},
	}
	tests = append(tests, Test{input, database})

	// Test adding hstore extension.
	input = `
extension "hstore" {}
	`
	database = &db.Database{
		Extensions: []*db.Extension{
			&db.Extension{
				ExtName: "hstore",
			},
		},
	}
	tests = append(tests, Test{input, database})

	for _, test := range tests {
		node, err := ParseBytes([]byte(test.Input))
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(node, test.Database) {
			t.Errorf("ParseBytes failure")
			spew.Dump(node, test.Database)
		}
	}
}

func TestConvertDatabase(t *testing.T) {
	type Test struct {
		Value    Database
		Database *db.Database
	}
	tests := []Test{
		Test{
			Database{},
			&db.Database{},
		},
		Test{
			Database{
				Schema: map[string]Schema{
					"public": Schema{},
				},
			},
			&db.Database{
				Schemas: []*db.Schema{
					&db.Schema{
						SchemaName: "public",
					},
				},
			},
		},
	}
	for _, test := range tests {
		node := convertDatabase(test.Value)
		if !reflect.DeepEqual(node, test.Database) {
			t.Errorf("convertDatabase failure")
			spew.Dump(node, test.Database)
		}
	}
}

func TestConvertSchema(t *testing.T) {
	type Test struct {
		SchemaName string
		Value      Schema
		Schema     *db.Schema
	}
	tests := []Test{
		Test{
			"public",
			Schema{},
			&db.Schema{
				SchemaName: "public",
			},
		},
		Test{
			"public",
			Schema{
				Table: map[string]Table{
					"users": Table{},
				},
			},
			&db.Schema{
				SchemaName: "public",
				Tables: []*db.Table{
					&db.Table{
						TableName: "users",
					},
				},
			},
		},
	}
	for _, test := range tests {
		node := convertSchema(test.SchemaName, test.Value)
		if !reflect.DeepEqual(node, test.Schema) {
			t.Errorf("convertSchema failure")
			spew.Dump(node, test.Schema)
		}
	}

}
