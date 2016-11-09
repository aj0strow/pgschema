package ab

import (
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func TestMatchExtensions(t *testing.T) {
	type Test struct {
		Name    string
		A       []db.ExtensionNode
		B       []db.ExtensionNode
		Matches []ExtensionMatch
	}
	tests := []Test{
		Test{
			"empty extension list",
			nil,
			nil,
			nil,
		},
		Test{
			"add uuid extension",
			[]db.ExtensionNode{
				db.ExtensionNode{
					Extension: db.Extension{
						ExtName: "uuid-ossp",
					},
				},
			},
			nil,
			[]ExtensionMatch{
				ExtensionMatch{
					A: &db.Extension{
						ExtName: "uuid-ossp",
					},
				},
			},
		},
		Test{
			"plpgsql already exists",
			nil,
			[]db.ExtensionNode{
				db.ExtensionNode{
					Extension: db.Extension{
						ExtName: "plpgsql",
					},
				},
			},
			[]ExtensionMatch{
				ExtensionMatch{
					A: nil,
					B: &db.Extension{
						ExtName: "plpgsql",
					},
				},
			},
		},
		Test{
			"citext already exists",
			[]db.ExtensionNode{
				db.ExtensionNode{
					Extension: db.Extension{
						ExtName: "citext",
					},
				},
			},
			[]db.ExtensionNode{
				db.ExtensionNode{
					Extension: db.Extension{
						ExtName: "citext",
					},
				},
			},
			[]ExtensionMatch{
				ExtensionMatch{
					A: &db.Extension{
						ExtName: "citext",
					},
					B: &db.Extension{
						ExtName: "citext",
					},
				},
			},
		},
	}
	for _, test := range tests {
		matches := MatchExtensions(test.A, test.B)
		if !reflect.DeepEqual(matches, test.Matches) {
			t.Errorf("MatchExtensions => %s", test.Name)
		}
	}
}
