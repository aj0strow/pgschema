package plan

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

func TestIndexChanges(t *testing.T) {
	type Test struct {
		Name       string
		IndexMatch ab.IndexMatch
		Changes    []Change
	}
	tests := []Test{
		Test{
			"drop existing index",
			ab.IndexMatch{
				B: &db.Index{
					IndexName: "users_pkey",
				},
			},
			[]Change{
				DropIndex{"users_pkey"},
			},
		},
	}
	for _, test := range tests {
		changes := IndexChanges(test.IndexMatch)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("IndexChanges => %s", test.Name)
		}
	}
}
