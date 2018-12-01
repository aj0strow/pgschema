package ab

import (
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func TestFindColumNode(t *testing.T) {
	type Test struct {
		Name       string
		Columns    []*db.Column
		SearchName string
		Found      *db.Column
	}
	tests := []Test{
		Test{
			`nil column node list`,
			nil,
			"test1",
			nil,
		},
		Test{
			`empty column list`,
			[]*db.Column{},
			"test1",
			nil,
		},
		Test{
			`wrong name`,
			[]*db.Column{
				&db.Column{ColumnName: "test2"},
			},
			"test1",
			nil,
		},
		Test{
			`correct name`,
			[]*db.Column{
				&db.Column{ColumnName: "test1"},
				&db.Column{ColumnName: "test2"},
			},
			"test2",
			&db.Column{ColumnName: "test2"},
		},
	}
	for _, test := range tests {
		node := findColumn(test.Columns, test.SearchName)
		if !reflect.DeepEqual(node, test.Found) {
			t.Errorf("findColumnNode - %s", test.Name)
		}
	}
}

func TestMatchColumnNodes(t *testing.T) {
	type Test struct {
		Name    string
		A       []*db.Column
		B       []*db.Column
		Matches []ColumnMatch
	}
	tests := []Test{
		Test{
			"nil column node lists",
			nil,
			nil,
			nil,
		},
		Test{
			"empty column node lists",
			[]*db.Column{},
			[]*db.Column{},
			nil,
		},
		Test{
			"columns in a only",
			[]*db.Column{
				&db.Column{ColumnName: "email"},
			},
			nil,
			[]ColumnMatch{
				ColumnMatch{
					A: &db.Column{ColumnName: "email"},
					B: nil,
				},
			},
		},
		Test{
			"columns in b only",
			nil,
			[]*db.Column{
				&db.Column{ColumnName: "dob"},
			},
			[]ColumnMatch{
				ColumnMatch{
					A: nil,
					B: &db.Column{ColumnName: "dob"},
				},
			},
		},
		Test{
			"multiple columns",
			[]*db.Column{
				&db.Column{ColumnName: "one"},
				&db.Column{ColumnName: "two"},
			},
			nil,
			[]ColumnMatch{
				ColumnMatch{
					A: &db.Column{ColumnName: "one"},
					B: nil,
				},
				ColumnMatch{
					A: &db.Column{ColumnName: "two"},
					B: nil,
				},
			},
		},
	}
	for _, test := range tests {
		matches := MatchColumns(test.A, test.B)
		if !reflect.DeepEqual(matches, test.Matches) {
			t.Errorf("MatchColumnNodes => %s", test.Name)
		}
	}
}
