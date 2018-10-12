package sql

import (
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	cs := []struct {
		Name         string
		SQL          string
		DatabaseNode db.DatabaseNode
	}{
		{
			`empty file`,
			``,
			db.DatabaseNode{},
		},
		{
			"parse extensions",
			`EXTENSION citext;`,
			db.DatabaseNode{
				ExtensionNodes: []db.ExtensionNode{
					db.ExtensionNode{
						Extension: db.Extension{
							ExtName: "citext",
						},
					},
				},
			},
		},
		{
			"parse schema declaration",
			"SCHEMA default;",
			db.DatabaseNode{
				SchemaNodes: []db.SchemaNode{
					db.SchemaNode{
						Schema: db.Schema{
							SchemaName: "default",
						},
					},
				},
			},
		},
		{
			"parse table definitions",
			"SCHEMA default; TABLE users (id serial);",
			db.DatabaseNode{
				SchemaNodes: []db.SchemaNode{
					db.SchemaNode{
						Schema: db.Schema{
							SchemaName: "default",
						},
						TableNodes: []db.TableNode{
							db.TableNode{
								Table: db.Table{
									TableName: "users",
								},
								ColumnNodes: []db.ColumnNode{
									db.ColumnNode{
										Column: db.Column{
											ColumnName: "id",
											DataType:   "serial",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, c := range cs {
		n, err := parse(c.SQL)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(n, c.DatabaseNode) {
			t.Errorf(c.Name)
			t.Errorf("%+v", n)
		}
	}
}

func TestParseColumn(t *testing.T) {
	cs := []struct {
		Fragment string
		Column   db.Column
		Indexes  []db.Index
	}{
		{
			`id serial`,
			db.Column{
				ColumnName: "id",
				DataType:   "serial",
			},
			nil,
		},
		{
			`title text NOT NULL`,
			db.Column{
				ColumnName: "title",
				DataType:   "text",
				NotNull:    true,
			},
			nil,
		},
		{
			`id serial PRIMARY KEY`,
			db.Column{
				ColumnName: "id",
				DataType:   "serial",
			},
			[]db.Index{
				db.Index{
					TableName: "posts",
					IndexName: "posts_pkey",
					Exprs:     []string{"id"},
					Primary:   true,
					Unique:    true,
				},
			},
		},
		{
			`name citext UNIQUE`,
			db.Column{
				ColumnName: "name",
				DataType:   "citext",
			},
			[]db.Index{
				db.Index{
					TableName: "posts",
					IndexName: "posts_name_key",
					Exprs:     []string{"name"},
					Unique:    true,
				},
			},
		},
		{
			`title text DEFAULT ''`,
			db.Column{
				ColumnName: "title",
				DataType:   "text",
				Default:    "''",
			},
			nil,
		},
		{
			`created_at timestamptz DEFAULT now()`,
			db.Column{
				ColumnName: "created_at",
				DataType:   "timestamptz",
				Default:    "now()",
			},
			nil,
		},
		{
			`views bigint DEFAULT 0`,
			db.Column{
				ColumnName: "views",
				DataType:   "bigint",
				Default:    "0",
			},
			nil,
		},
	}
	for _, c := range cs {
		p := newParser(lex(c.Fragment))
		tableNode := db.TableNode{
			Table: db.Table{
				TableName: "posts",
			},
		}
		columnNode := db.ColumnNode{}
		err := parseColumnNode(p, &tableNode, &columnNode)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(columnNode.Column, c.Column) {
			t.Errorf(c.Fragment)
			t.Errorf("Column: %+v", columnNode.Column)
		}
		for i, idx := range c.Indexes {
			if !reflect.DeepEqual(tableNode.IndexNodes[i].Index, idx) {
				t.Errorf(c.Fragment)
				t.Errorf("Index: %+v", idx)
			}
		}
	}
}
