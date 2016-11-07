package pgschema

import (
	"github.com/aj0strow/pgschema/info"
	"reflect"
	"testing"
)

func TestPlanTableMatch(t *testing.T) {
	type Test struct {
		Name       string
		TableMatch TableMatch
		Changes    []Change
	}
	tests := []Test{
		Test{
			"create new table",
			TableMatch{
				A: &info.Table{
					TableName: "users",
				},
				B: nil,
			},
			[]Change{
				CreateTable{"users"},
			},
		},
		Test{
			"drop old table",
			TableMatch{
				A: nil,
				B: &info.Table{
					TableName: "customers",
				},
			},
			[]Change{
				DropTable{"customers"},
			},
		},
	}
	for _, test := range tests {
		changes := planTableMatch(test.TableMatch)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("planTableMatch => %s", test.Name)
		}
	}
}

func TestPlanColumnMatch(t *testing.T) {
	type Test struct {
		Name        string
		ColumnMatch ColumnMatch
		Changes     []Change
	}
	tests := []Test{
		Test{
			"drop old column",
			ColumnMatch{
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
			ColumnMatch{
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
			ColumnMatch{
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
		changes := planColumnMatch(test.ColumnMatch)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("planColumnMatch => %s", test.Name)
		}
	}
}
