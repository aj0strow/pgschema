package ab

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
)

func TestMatchIndexes(t *testing.T) {
	type Test struct {
		Name    string
		A       []*db.Index
		B       []*db.Index
		Matches []IndexMatch
	}
	tests := []Test{
		Test{
			"empty index list",
			nil,
			nil,
			nil,
		},
		Test{
			"add new index",
			[]*db.Index{
				&db.Index{
					IndexName: "users_pkey",
				},
			},
			nil,
			[]IndexMatch{
				IndexMatch{
					A: &db.Index{
						IndexName: "users_pkey",
					},
				},
			},
		},
		Test{
			"remove existing index",
			nil,
			[]*db.Index{
				&db.Index{
					IndexName: "users_email_idx",
				},
			},
			[]IndexMatch{
				IndexMatch{
					B: &db.Index{
						IndexName: "users_email_idx",
					},
				},
			},
		},
		Test{
			"existing index noop",
			[]*db.Index{
				&db.Index{
					IndexName: "users_nickname_key",
				},
			},
			[]*db.Index{
				&db.Index{
					IndexName: "users_nickname_key",
				},
			},
			[]IndexMatch{
				IndexMatch{
					A: &db.Index{
						IndexName: "users_nickname_key",
					},
					B: &db.Index{
						IndexName: "users_nickname_key",
					},
				},
			},
		},
	}
	for _, test := range tests {
		matches := MatchIndexes(test.A, test.B)
		if !reflect.DeepEqual(matches, test.Matches) {
			t.Errorf("MatchIndexes => %s", test.Name)
		}
	}
}
