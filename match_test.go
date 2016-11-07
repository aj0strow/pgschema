package pgschema

import (
	"reflect"
	"testing"
)

func newColumnNode(name string) ColumnNode {
	return ColumnNode{
		Column{
			ColumnName: name,
		},
	}
}

func ptrColumnNode(name string) *ColumnNode {
	node := newColumnNode(name)
	return &node
}

func ptrColumn(name string) *Column {
	return &Column{
		ColumnName: name,
	}
}

func TestFindColumNode(t *testing.T) {
	type Test struct {
		Name        string
		ColumnNodes []ColumnNode
		SearchName  string
		Found       *ColumnNode
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
			[]ColumnNode{},
			"test1",
			nil,
		},
		Test{
			`wrong name`,
			[]ColumnNode{
				newColumnNode("test2"),
			},
			"test1",
			nil,
		},
		Test{
			`correct name`,
			[]ColumnNode{
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
		A       []ColumnNode
		B       []ColumnNode
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
			[]ColumnNode{},
			[]ColumnNode{},
			nil,
		},
		Test{
			"columns in a only",
			[]ColumnNode{
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
			[]ColumnNode{
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
			[]ColumnNode{
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
		matches := MatchColumnNodes(test.A, test.B)
		if !reflect.DeepEqual(matches, test.Matches) {
			t.Errorf("MatchColumnNodes => %s", test.Name)
		}
	}
}
