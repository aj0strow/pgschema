package hcl

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestParseBytes(t *testing.T) {
	type Test struct {
		Input        string
		DatabaseNode db.DatabaseNode
	}
	var (
		tests    []Test
		input    string
		database db.DatabaseNode
		tables   []db.TableNode
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
	tables = []db.TableNode{
		db.TableNode{
			Table: db.Table{
				TableName: "users",
			},
			ColumnNodes: []db.ColumnNode{
				db.ColumnNode{
					Column: db.Column{
						ColumnName: "email",
						DataType:   "text",
					},
				},
			},
		},
	}
	database = db.DatabaseNode{
		SchemaNodes: []db.SchemaNode{
			db.SchemaNode{
				Schema: db.Schema{
					SchemaName: "public",
				},
				TableNodes: tables,
			},
		},
	}
	tests = append(tests, Test{input, database})

	// Test adding hstore extension.
	input = `
extension "hstore" {}
	`
	database = db.DatabaseNode{
		ExtensionNodes: []db.ExtensionNode{
			db.ExtensionNode{
				db.Extension{
					ExtName: "hstore",
				},
			},
		},
	}
	tests = append(tests, Test{input, database})

	for _, test := range tests {
		node, err := ParseBytes([]byte(test.Input))
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(node, test.DatabaseNode) {
			t.Errorf("ParseBytes failure")
			spew.Dump(node, test.DatabaseNode)
		}
	}
}

func TestConvertDatabase(t *testing.T) {
	type Test struct {
		Value        Database
		DatabaseNode db.DatabaseNode
	}
	tests := []Test{
		Test{
			Database{},
			db.DatabaseNode{},
		},
		Test{
			Database{
				Schema: map[string]Schema{
					"public": Schema{},
				},
			},
			db.DatabaseNode{
				SchemaNodes: []db.SchemaNode{
					db.SchemaNode{
						Schema: db.Schema{
							SchemaName: "public",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		node := convertDatabase(test.Value)
		if !reflect.DeepEqual(node, test.DatabaseNode) {
			t.Errorf("convertDatabase failure")
			spew.Dump(node, test.DatabaseNode)
		}
	}
}

func TestConvertSchema(t *testing.T) {
	type Test struct {
		Key        string
		Value      Schema
		SchemaNode db.SchemaNode
	}
	tests := []Test{
		Test{
			"public",
			Schema{},
			db.SchemaNode{
				Schema: db.Schema{
					SchemaName: "public",
				},
			},
		},
		Test{
			"public",
			Schema{
				Table: map[string]Table{
					"users": Table{},
				},
			},
			db.SchemaNode{
				Schema: db.Schema{
					SchemaName: "public",
				},
				TableNodes: []db.TableNode{
					db.TableNode{
						Table: db.Table{
							TableName: "users",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		node := convertSchema(test.Key, test.Value)
		if !reflect.DeepEqual(node, test.SchemaNode) {
			t.Errorf("convertSchema failure")
			spew.Dump(node, test.SchemaNode)
		}
	}

}

func TestConvertTable(t *testing.T) {
	type Test struct {
		Key       string
		Value     Table
		TableNode db.TableNode
	}
	tests := []Test{
		Test{
			"users",
			Table{},
			db.TableNode{
				Table: db.Table{
					TableName: "users",
				},
			},
		},
		Test{
			"customers",
			Table{
				Column: map[string]Column{
					"email": Column{
						Type: "text",
					},
				},
			},
			db.TableNode{
				Table: db.Table{
					TableName: "customers",
				},
				ColumnNodes: []db.ColumnNode{
					db.ColumnNode{
						Column: db.Column{
							ColumnName: "email",
							DataType:   "text",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		node := convertTable(test.Key, test.Value)
		if !reflect.DeepEqual(node, test.TableNode) {
			t.Errorf("convertTable failure")
			t.Errorf("have: %#v\nwant: %#v\n", node, test.TableNode)
		}
	}
}

func TestConvertColumn(t *testing.T) {
	type Test struct {
		Key        string
		Value      Column
		ColumnNode db.ColumnNode
	}
	tests := []Test{
		Test{
			"email",
			Column{
				Type: "text",
			},
			db.ColumnNode{
				Column: db.Column{
					ColumnName: "email",
					DataType:   "text",
				},
			},
		},
	}
	for _, test := range tests {
		node := convertColumn(test.Key, test.Value)
		if !reflect.DeepEqual(node, test.ColumnNode) {
			t.Errorf("convertColumn failure")
			t.Errorf("have: %#v\nwant: %#v\n", node, test.ColumnNode)
		}
	}
}
