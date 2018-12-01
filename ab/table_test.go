package ab

import (
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func TestMatchTableNodes(t *testing.T) {
	type Test struct {
		Name    string
		A       []*db.Table
		B       []*db.Table
		Matches []TableMatch
	}
	tests := []Test{
		Test{
			"multiple tables",
			[]*db.Table{
				&db.Table{TableName: "users"},
				&db.Table{TableName: "passwords"},
			},
			nil,
			[]TableMatch{
				TableMatch{
					A: &db.Table{TableName: "users"},
					B: nil,
				},
				TableMatch{
					A: &db.Table{TableName: "passwords"},
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
