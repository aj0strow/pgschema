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
		Test{
			`ignore old indexes`,
			[]ab.IndexMatch{
				ab.IndexMatch{
					B: &db.Index{},
				},
			},
			nil,
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

func TestDropIndexes(t *testing.T) {
	type Test struct {
		Name         string
		IndexMatches []ab.IndexMatch
		DropIndexes  []DropIndex
	}
	tests := []Test{
		Test{
			`empty index list`,
			nil,
			nil,
		},
		Test{
			`ignore new indexes`,
			[]ab.IndexMatch{
				ab.IndexMatch{
					A: &db.Index{},
				},
			},
			nil,
		},
		Test{
			`drop old indexes`,
			[]ab.IndexMatch{
				ab.IndexMatch{
					B: &db.Index{},
				},
			},
			[]DropIndex{
				DropIndex{
					Index: &db.Index{},
				},
			},
		},
	}
	for _, test := range tests {
		xs := dropIndexes(test.IndexMatches)
		if !reflect.DeepEqual(xs, test.DropIndexes) {
			t.Errorf("dropIndexes => %s", test.Name)
			spew.Dump(xs, test.DropIndexes)
		}
	}
}
