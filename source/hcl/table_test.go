package hcl

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestConvertTable(t *testing.T) {
	type Test struct {
		TableName string
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
		Test{
			"customers",
			Table{
				Column: map[string]Column{
					"id": Column{
						Type:       "text",
						PrimaryKey: true,
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
							ColumnName: "id",
							DataType:   "text",
							NotNull:    true,
						},
					},
				},
				IndexNodes: []db.IndexNode{
					db.IndexNode{
						Index: db.Index{
							TableName: "customers",
							IndexName: "customers_pkey",
							Exprs:     []string{"id"},
							Unique:    true,
							Primary:   true,
						},
					},
				},
			},
		},
		Test{
			"events",
			Table{
				PrimaryKey: []string{"source", "time"},
			},
			db.TableNode{
				Table: db.Table{
					TableName: "events",
				},
				IndexNodes: []db.IndexNode{
					db.IndexNode{
						Index: db.Index{
							TableName: "events",
							IndexName: "events_pkey",
							Exprs:     []string{"source", "time"},
							Unique:    true,
							Primary:   true,
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		node := convertTable(test.TableName, test.Value)
		if !reflect.DeepEqual(node, test.TableNode) {
			t.Errorf("convertTable failure")
			spew.Dump(node, test.TableNode)
		}
	}
}
