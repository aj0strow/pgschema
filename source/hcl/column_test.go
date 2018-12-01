package hcl

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestConvertColumn(t *testing.T) {
	type Test struct {
		Key    string
		Value  Column
		Column *db.Column
	}
	tests := []Test{
		Test{
			"name",
			Column{
				Type: "text",
			},
			&db.Column{
				ColumnName: "name",
				DataType:   "text",
			},
		},
		Test{
			"email",
			Column{
				Type:    "text",
				NotNull: true,
			},
			&db.Column{
				ColumnName: "email",
				DataType:   "text",
				NotNull:    true,
			},
		},
		Test{
			"created_at",
			Column{
				Type:    "timestamptz",
				Default: "now()",
			},
			&db.Column{
				ColumnName: "created_at",
				DataType:   "timestamptz",
				Default:    "now()",
			},
		},
		Test{
			"amount",
			Column{
				Type: "numeric(11,2)",
			},
			&db.Column{
				ColumnName:       "amount",
				DataType:         "numeric",
				NumericPrecision: 11,
				NumericScale:     2,
			},
		},
		Test{
			"tick_size",
			Column{
				Type: "numeric(24, 16)",
			},
			&db.Column{
				ColumnName:       "tick_size",
				DataType:         "numeric",
				NumericPrecision: 24,
				NumericScale:     16,
			},
		},
		Test{
			"upper4",
			Column{
				Type: "integer[4]",
			},
			&db.Column{
				ColumnName: "upper4",
				DataType:   "integer",
				Array:      true,
			},
		},
	}
	for _, test := range tests {
		node := convertColumn(test.Key, test.Value)
		if !reflect.DeepEqual(node, test.Column) {
			t.Errorf("convertColumn failure")
			spew.Dump(node, test.Column)
		}
	}
}
