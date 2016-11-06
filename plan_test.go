package pgschema

import (
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
				A: &Table{
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
				B: &Table{
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
				B: &Column{
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
				A: &Column{
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
				A: &Column{
					ColumnName: "email",
					DataType:   "text",
				},
				B: &Column{
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
