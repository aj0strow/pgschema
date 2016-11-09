package ab

import (
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func newTable(name string) db.Table {
	return db.Table{
		TableName: name,
	}
}

func newTableNode(name string) db.TableNode {
	return db.TableNode{
		Table: newTable(name),
	}
}

func ptrTable(name string) *db.Table {
	table := newTable(name)
	return &table
}

func TestMatchTableNodes(t *testing.T) {
	type Test struct {
		Name    string
		A       []db.TableNode
		B       []db.TableNode
		Matches []TableMatch
	}
	tests := []Test{
		Test{
			"multiple tables",
			[]db.TableNode{
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
