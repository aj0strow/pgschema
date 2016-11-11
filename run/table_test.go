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
					A: &db.Table{},
				},
			},
			[]CreateTable{
				CreateTable{
					Table: &db.Table{},
				},
			},
		},
		Test{
			`ignore existing table`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
					B: &db.Table{},
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

func TestDropTables(t *testing.T) {
	type Test struct {
		Name         string
		TableMatches []ab.TableMatch
		DropTables   []DropTable
	}
	tests := []Test{
		Test{
			`empty table list`,
			nil,
			nil,
		},
		Test{
			`drop existing table`,
			[]ab.TableMatch{
				ab.TableMatch{
					B: &db.Table{},
				},
			},
			[]DropTable{
				DropTable{
					Table: &db.Table{},
				},
			},
		},
		Test{
			`ignore required table`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
					B: &db.Table{},
				},
			},
			nil,
		},
	}
	for _, test := range tests {
		xs := dropTables(test.TableMatches)
		if !reflect.DeepEqual(xs, test.DropTables) {
			t.Errorf("dropTables => %s", test.Name)
			spew.Dump(xs, test.DropTables)
		}
	}
}
