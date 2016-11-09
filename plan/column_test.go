package plan

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func TestColumnChanges(t *testing.T) {
	type Test struct {
		Name        string
		ColumnMatch ab.ColumnMatch
		Changes     []Change
	}
	tests := []Test{
		Test{
			"drop old column",
			ab.ColumnMatch{
				A: nil,
				B: &db.Column{
					ColumnName: "email",
				},
			},
			[]Change{
				DropColumn{"email"},
			},
		},
		Test{
			"add new column",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName: "email",
					DataType:   "citext",
				},
				B: nil,
			},
			[]Change{
				AddColumn{"email", "citext"},
			},
		},
		Test{
			"change column type",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName: "email",
					DataType:   "text",
				},
				B: &db.Column{
					ColumnName: "email",
					DataType:   "citext",
				},
			},
			[]Change{
				AlterColumn{
					"email",
					SetDataType{"text"},
				},
			},
		},
	}
	for _, test := range tests {
		changes := ColumnChanges(test.ColumnMatch)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("planColumnMatch => %s", test.Name)
		}
	}
}
