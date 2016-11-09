package plan

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

func TestExtensionChanges(t *testing.T) {
	type Test struct {
		Name    string
		Match   ab.ExtensionMatch
		Changes []Change
	}
	tests := []Test{
		Test{
			"ignore plpgsql",
			ab.ExtensionMatch{
				A: nil,
				B: &db.Extension{
					ExtName: "plpgsql",
				},
			},
			nil,
		},
		Test{
			"create uuid extension",
			ab.ExtensionMatch{
				A: &db.Extension{
					ExtName: "uuid-ossp",
				},
			},
			[]Change{
				CreateExtension{ExtName: "uuid-ossp"},
			},
		},
		Test{
			"drop citext extension",
			ab.ExtensionMatch{
				A: nil,
				B: &db.Extension{
					ExtName: "citext",
				},
			},
			[]Change{
				DropExtension{ExtName: "citext"},
			},
		},
		Test{
			"extension already exists",
			ab.ExtensionMatch{
				A: &db.Extension{
					ExtName: "hstore",
				},
				B: &db.Extension{
					ExtName: "hstore",
				},
			},
			nil,
		},
	}
	for _, test := range tests {
		changes := ExtensionChanges(test.Match)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("ExtensionChanges => %s", test.Name)
		}
	}
}
