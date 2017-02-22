package hcl

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestConvertColumn(t *testing.T) {
	type Test struct {
		Key        string
		Value      Column
		ColumnNode db.ColumnNode
	}
	tests := []Test{
		Test{
			"name",
			Column{
				Type: "text",
			},
			db.ColumnNode{
				Column: db.Column{
					ColumnName: "name",
					DataType:   "text",
				},
			},
		},
		Test{
			"email",
			Column{
				Type:    "text",
				NotNull: true,
			},
			db.ColumnNode{
				Column: db.Column{
					ColumnName: "email",
					DataType:   "text",
					NotNull:    true,
				},
			},
		},
		Test{
			"created_at",
			Column{
				Type:    "timestamptz",
				Default: "now()",
			},
			db.ColumnNode{
				Column: db.Column{
					ColumnName: "created_at",
					DataType:   "timestamptz",
					Default:    "now()",
				},
			},
		},
		Test{
			"amount",
			Column{
				Type: "numeric(11,2)",
			},
			db.ColumnNode{
				Column: db.Column{
					ColumnName:       "amount",
					DataType:         "numeric",
					NumericPrecision: 11,
					NumericScale:     2,
				},
			},
		},
		Test{
			"upper4",
			Column{
				Type: "integer[4]",
			},
			db.ColumnNode{
				Column: db.Column{
					ColumnName: "upper4",
					DataType:   "integer",
					Array:      true,
				},
			},
		},
	}
	for _, test := range tests {
		node := convertColumn(test.Key, test.Value)
		if !reflect.DeepEqual(node, test.ColumnNode) {
			t.Errorf("convertColumn failure")
			spew.Dump(node, test.ColumnNode)
		}
	}
}
