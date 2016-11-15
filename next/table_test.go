package next

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
		Test{
			`add columns`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
					ColumnMatches: []ab.ColumnMatch{
						ab.ColumnMatch{
							A: &db.Column{},
						},
					},
				},
			},
			[]CreateTable{
				CreateTable{
					Table: &db.Table{},
					AddColumns: []AddColumn{
						AddColumn{
							Column: &db.Column{},
						},
					},
				},
			},
		},
		Test{
			`create indexes`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
					IndexMatches: []ab.IndexMatch{
						ab.IndexMatch{
							A: &db.Index{},
						},
					},
				},
			},
			[]CreateTable{
				CreateTable{
					Table: &db.Table{},
					CreateIndexes: []CreateIndex{
						CreateIndex{
							Index: &db.Index{},
						},
					},
				},
			},
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

func TestAlterTables(t *testing.T) {
	type Test struct {
		Name         string
		TableMatches []ab.TableMatch
		AlterTables  []AlterTable
	}
	tests := []Test{
		Test{
			`empty table matches`,
			nil,
			nil,
		},
		Test{
			`ignore new tables`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
				},
			},
			nil,
		},
		Test{
			`ignore old tables`,
			[]ab.TableMatch{
				ab.TableMatch{
					B: &db.Table{},
				},
			},
			nil,
		},
		Test{
			`add columns`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
					B: &db.Table{},
					ColumnMatches: []ab.ColumnMatch{
						ab.ColumnMatch{
							A: &db.Column{},
						},
					},
				},
			},
			[]AlterTable{
				AlterTable{
					Table: &db.Table{},
					AddColumns: []AddColumn{
						AddColumn{
							Column: &db.Column{},
						},
					},
				},
			},
		},
		Test{
			`drop columns`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
					B: &db.Table{},
					ColumnMatches: []ab.ColumnMatch{
						ab.ColumnMatch{
							B: &db.Column{},
						},
					},
				},
			},
			[]AlterTable{
				AlterTable{
					Table: &db.Table{},
					DropColumns: []DropColumn{
						DropColumn{
							Column: &db.Column{},
						},
					},
				},
			},
		},
		Test{
			`create new indexes`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
					B: &db.Table{},
					IndexMatches: []ab.IndexMatch{
						ab.IndexMatch{
							A: &db.Index{},
						},
					},
				},
			},
			[]AlterTable{
				AlterTable{
					Table: &db.Table{},
					CreateIndexes: []CreateIndex{
						CreateIndex{
							Index: &db.Index{},
						},
					},
				},
			},
		},
		Test{
			`drop old indexes`,
			[]ab.TableMatch{
				ab.TableMatch{
					A: &db.Table{},
					B: &db.Table{},
					IndexMatches: []ab.IndexMatch{
						ab.IndexMatch{
							B: &db.Index{},
						},
					},
				},
			},
			[]AlterTable{
				AlterTable{
					Table: &db.Table{},
					DropIndexes: []DropIndex{
						DropIndex{
							Index: &db.Index{},
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		xs := alterTables(test.TableMatches)
		if !reflect.DeepEqual(xs, test.AlterTables) {
			t.Errorf("alterTables => %s", test.Name)
			spew.Dump(xs, test.AlterTables)
		}
	}
}
