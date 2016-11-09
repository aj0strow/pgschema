package ab

import (
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func newColumnNode(name string) db.ColumnNode {
	return db.ColumnNode{
		db.Column{
			ColumnName: name,
		},
	}
}

func ptrColumnNode(name string) *db.ColumnNode {
	node := newColumnNode(name)
	return &node
}

func ptrColumn(name string) *db.Column {
	return &db.Column{
		ColumnName: name,
	}
}

func TestFindColumNode(t *testing.T) {
	type Test struct {
		Name        string
		ColumnNodes []db.ColumnNode
		SearchName  string
		Found       *db.ColumnNode
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
			[]db.ColumnNode{},
			"test1",
			nil,
		},
		Test{
			`wrong name`,
			[]db.ColumnNode{
				newColumnNode("test2"),
			},
			"test1",
			nil,
		},
		Test{
			`correct name`,
			[]db.ColumnNode{
				newColumnNode("test1"),
				newColumnNode("test2"),
			},
			"test2",
			ptrColumnNode("test2"),
		},
	}
	for _, test := range tests {
		node := findColumnNode(test.ColumnNodes, test.SearchName)
		if !reflect.DeepEqual(node, test.Found) {
			t.Errorf("findColumnNode - %s", test.Name)
		}
	}
}

func TestMatchColumnNodes(t *testing.T) {
	type Test struct {
		Name    string
		A       []db.ColumnNode
		B       []db.ColumnNode
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
			[]db.ColumnNode{},
			[]db.ColumnNode{},
			nil,
		},
		Test{
			"columns in a only",
			[]db.ColumnNode{
				newColumnNode("email"),
			},
			nil,
			[]ColumnMatch{
				ColumnMatch{
					A: ptrColumn("email"),
					B: nil,
				},
			},
		},
		Test{
			"columns in b only",
			nil,
			[]db.ColumnNode{
				newColumnNode("dob"),
			},
			[]ColumnMatch{
				ColumnMatch{
					A: nil,
					B: ptrColumn("dob"),
				},
			},
		},
		Test{
			"multiple columns",
			[]db.ColumnNode{
				newColumnNode("one"),
				newColumnNode("two"),
			},
			nil,
			[]ColumnMatch{
				ColumnMatch{
					A: ptrColumn("one"),
					B: nil,
				},
				ColumnMatch{
					A: ptrColumn("two"),
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
