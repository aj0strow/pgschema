package run

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestCreateTables(t *testing.T) {
	type Test struct {
		Name         string
		TableMatches []ab.TableMatch
		CreateTables []CreateTable
	}
	tests := []Test{
		Test{
			`empty table list`,
			nil,
			nil,
		},
		Test{
			`create required table`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{
						TableName: "users",
					},
				},
			},
			[]CreateTable{
				CreateTable{
					Table: &db.Table{
						TableName: "users",
					},
				},
			},
		},
		Test{
			`ignore existing table`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{
						TableName: "users",
					},
					B: &db.Table{
						TableName: "users",
					},
				},
			},
			nil,
		},
	}
	for _, test := range tests {
		xs := createTables(test.TableMatches)
		if !reflect.DeepEqual(xs, test.CreateTables) {
			t.Errorf("createTables => %s", test.Name)
			spew.Dump(xs, test.CreateTables)
		}
	}
}
