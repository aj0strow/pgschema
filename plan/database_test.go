package plan

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestUpdateDatabase(t *testing.T) {
	type Test struct {
		Name           string
		DatabaseMatch  ab.DatabaseMatch
		UpdateDatabase UpdateDatabase
	}
	tests := []Test{
		Test{
			`empty database`,
			ab.DatabaseMatch{},
			UpdateDatabase{},
		},
		Test{
			`create schemas`,
			ab.DatabaseMatch{
				SchemaMatches: []ab.SchemaMatch{
					ab.SchemaMatch{
						A: &db.Schema{},
					},
				},
			},
			UpdateDatabase{
				CreateSchemas: []CreateSchema{
					CreateSchema{
						Schema: &db.Schema{},
					},
				},
			},
		},
		Test{
			`create extensions`,
			ab.DatabaseMatch{
				ExtensionMatches: []ab.ExtensionMatch{
					ab.ExtensionMatch{
						A: &db.Extension{},
					},
				},
			},
			UpdateDatabase{
				CreateExtensions: []CreateExtension{
					CreateExtension{
						Extension: &db.Extension{},
					},
				},
			},
		},
	}
	for _, test := range tests {
		x := updateDatabase(test.DatabaseMatch)
		if !reflect.DeepEqual(x, test.UpdateDatabase) {
			t.Errorf("updateDatabase => %s", test.Name)
			spew.Dump(x, test.UpdateDatabase)
		}
	}
}
