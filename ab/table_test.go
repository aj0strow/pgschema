package ab

import (
	"github.com/aj0strow/pgschema/info"
	"github.com/aj0strow/pgschema/tree"
	"reflect"
	"testing"
)

func newTable(name string) info.Table {
	return info.Table{
		TableName: name,
	}
}

func newTableNode(name string) tree.TableNode {
	return tree.TableNode{
		Table: newTable(name),
	}
}

func ptrTable(name string) *info.Table {
	table := newTable(name)
	return &table
}

func TestMatchTableNodes(t *testing.T) {
	type Test struct {
		Name    string
		A       []tree.TableNode
		B       []tree.TableNode
		Matches []TableMatch
	}
	tests := []Test{
		Test{
			"multiple tables",
			[]tree.TableNode{
				newTableNode("users"),
				newTableNode("passwords"),
			},
			nil,
			[]TableMatch{
				TableMatch{
					A: ptrTable("users"),
					B: nil,
				},
				TableMatch{
					A: ptrTable("passwords"),
					B: nil,
				},
			},
		},
	}
	for _, test := range tests {
		matches := MatchTables(test.A, test.B)
		if !reflect.DeepEqual(matches, test.Matches) {
			t.Errorf("MatchTableNodes => %s", test.Name)
		}
	}
}
