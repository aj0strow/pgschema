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
		Table     *db.Table
	}
	tests := []Test{
		Test{
			"users",
			Table{},
			&db.Table{
				TableName: "users",
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
			&db.Table{
				TableName: "customers",
				Columns: []*db.Column{
					&db.Column{
						ColumnName: "email",
						DataType:   "text",
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
			&db.Table{
				TableName: "customers",
				Columns: []*db.Column{
					&db.Column{
						ColumnName: "id",
						DataType:   "text",
						NotNull:    true,
					},
				},
				Indexes: []*db.Index{
					&db.Index{
						TableName: "customers",
						IndexName: "customers_pkey",
						Exprs:     []string{"id"},
						Unique:    true,
						Primary:   true,
					},
				},
			},
		},
		Test{
			"events",
			Table{
				PrimaryKey: []string{"source", "time"},
			},
			&db.Table{
				TableName: "events",
				Indexes: []*db.Index{
					&db.Index{
						TableName: "events",
						IndexName: "events_pkey",
						Exprs:     []string{"source", "time"},
						Unique:    true,
						Primary:   true,
					},
				},
			},
		},
	}
	for _, test := range tests {
		node := convertTable(test.TableName, test.Value)
		if !reflect.DeepEqual(node, test.Table) {
			t.Errorf("convertTable failure")
			spew.Dump(node, test.Table)
		}
	}
}
