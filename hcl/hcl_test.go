package hcl

import (
	"github.com/aj0strow/pgschema/info"
	"github.com/aj0strow/pgschema/tree"
	"reflect"
	"testing"
)

func TestParseBytes(t *testing.T) {
	type Test struct {
		Input        string
		DatabaseNode tree.DatabaseNode
	}
	var (
		tests    []Test
		input    string
		database tree.DatabaseNode
		tables   []tree.TableNode
	)
	input = `
schema "public" {
	table "users" {
		column "email" {
			type = "text"
		}
	}
}
	`
	tables = []tree.TableNode{
		tree.TableNode{
			Table: info.Table{
				TableName: "users",
			},
			ColumnNodes: []tree.ColumnNode{
				tree.ColumnNode{
					Column: info.Column{
						ColumnName: "email",
						DataType:   "text",
					},
				},
			},
		},
	}
	database = tree.DatabaseNode{
		SchemaNodes: []tree.SchemaNode{
			tree.SchemaNode{
				Schema: info.Schema{
					SchemaName: "public",
				},
				TableNodes: tables,
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
			t.Errorf("have: %#v\nwant: %#v\n", node, test.DatabaseNode)
		}
	}
}

func TestConvertDatabase(t *testing.T) {
	type Test struct {
		Value        Database
		DatabaseNode tree.DatabaseNode
	}
	tests := []Test{
		Test{
			Database{},
			tree.DatabaseNode{},
		},
		Test{
			Database{
				Schema: map[string]Schema{
					"public": Schema{},
				},
			},
			tree.DatabaseNode{
				SchemaNodes: []tree.SchemaNode{
					tree.SchemaNode{
						Schema: info.Schema{
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
			t.Errorf("have: %#v\nwant: %#v\n", node, test.DatabaseNode)
		}
	}
}

func TestConvertSchema(t *testing.T) {
	type Test struct {
		Key        string
		Value      Schema
		SchemaNode tree.SchemaNode
	}
	tests := []Test{
		Test{
			"public",
			Schema{},
			tree.SchemaNode{
				Schema: info.Schema{
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
			tree.SchemaNode{
				Schema: info.Schema{
					SchemaName: "public",
				},
				TableNodes: []tree.TableNode{
					tree.TableNode{
						Table: info.Table{
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
			t.Errorf("have: %#v\nwant: %#v\n", node, test.SchemaNode)
		}
	}

}

func TestConvertTable(t *testing.T) {
	type Test struct {
		Key       string
		Value     Table
		TableNode tree.TableNode
	}
	tests := []Test{
		Test{
			"users",
			Table{},
			tree.TableNode{
				Table: info.Table{
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
			tree.TableNode{
				Table: info.Table{
					TableName: "customers",
				},
				ColumnNodes: []tree.ColumnNode{
					tree.ColumnNode{
						Column: info.Column{
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
		ColumnNode tree.ColumnNode
	}
	tests := []Test{
		Test{
			"email",
			Column{
				Type: "text",
			},
			tree.ColumnNode{
				Column: info.Column{
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
