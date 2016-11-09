package plan

import (
	"github.com/aj0strow/pgschema/info"
	"github.com/aj0strow/pgschema/tree"
	"reflect"
	"testing"
)

func TestColumnChanges(t *testing.T) {
	type Test struct {
		Name        string
		ColumnMatch tree.ColumnMatch
		Changes     []Change
	}
	tests := []Test{
		Test{
			"drop old column",
			tree.ColumnMatch{
				A: nil,
				B: &info.Column{
					ColumnName: "email",
				},
			},
			[]Change{
				DropColumn{"email"},
			},
		},
		Test{
			"add new column",
			tree.ColumnMatch{
				A: &info.Column{
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
			tree.ColumnMatch{
				A: &info.Column{
					ColumnName: "email",
					DataType:   "text",
				},
				B: &info.Column{
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
