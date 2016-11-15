package next

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestCreateIndexes(t *testing.T) {
	type Test struct {
		Name          string
		IndexMatches  []ab.IndexMatch
		CreateIndexes []CreateIndex
	}
	tests := []Test{
		Test{
			`empty index list`,
			nil,
			nil,
		},
		Test{
			`create new index`,
			[]ab.IndexMatch{
				ab.IndexMatch{
					A: &db.Index{},
				},
			},
			[]CreateIndex{
				CreateIndex{
					Index: &db.Index{},
				},
			},
		},
	}
	for _, test := range tests {
		xs := createIndexes(test.IndexMatches)
		if !reflect.DeepEqual(xs, test.CreateIndexes) {
			t.Errorf("createIndexes => %s", test.Name)
			spew.Dump(xs, test.CreateIndexes)
		}
	}
}
