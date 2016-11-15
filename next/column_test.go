package next

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestAddColumns(t *testing.T) {
	type Test struct {
		Name          string
		ColumnMatches []ab.ColumnMatch
		AddColumns    []AddColumn
	}
	tests := []Test{
		Test{
			`empty column list`,
			nil,
			nil,
		},
		Test{
			`add required column`,
			[]ab.ColumnMatch{
				ab.ColumnMatch{
					A: &db.Column{},
				},
			},
			[]AddColumn{
				AddColumn{
					Column: &db.Column{},
				},
			},
		},
		Test{
			`ignore existing column`,
			[]ab.ColumnMatch{
				ab.ColumnMatch{
					A: &db.Column{},
					B: &db.Column{},
				},
			},
			nil,
		},
	}
	for _, test := range tests {
		xs := addColumns(test.ColumnMatches)
		if !reflect.DeepEqual(xs, test.AddColumns) {
			t.Errorf("addColumns => %s", test.Name)
			spew.Dump(xs, test.AddColumns)
		}
	}
}

func TestDropColumns(t *testing.T) {
	type Test struct {
		Name          string
		ColumnMatches []ab.ColumnMatch
		DropColumns   []DropColumn
	}
	tests := []Test{
		Test{
			`empty column list`,
			nil,
			nil,
		},
		Test{
			`drop existing column`,
			[]ab.ColumnMatch{
				ab.ColumnMatch{
					B: &db.Column{},
				},
			},
			[]DropColumn{
				DropColumn{
					Column: &db.Column{},
				},
			},
		},
	}
	for _, test := range tests {
		xs := dropColumns(test.ColumnMatches)
		if !reflect.DeepEqual(xs, test.DropColumns) {
			t.Errorf("dropColumns => %s", test.Name)
			spew.Dump(xs, test.DropColumns)
		}
	}
}

func TestAlterColumns(t *testing.T) {
	type Test struct {
		Name          string
		ColumnMatches []ab.ColumnMatch
		AlterColumns  []AlterColumn
	}
	tests := []Test{
		Test{
			`empty columns list`,
			nil,
			nil,
		},
		Test{
			`alter existing column`,
			[]ab.ColumnMatch{
				ab.ColumnMatch{
					A: &db.Column{},
					B: &db.Column{},
				},
			},
			[]AlterColumn{
				AlterColumn{
					Column: &db.Column{},
				},
			},
		},
	}
	for _, test := range tests {
		xs := alterColumns(test.ColumnMatches)
		if !reflect.DeepEqual(xs, test.AlterColumns) {
			t.Errorf("alterColumns => %s", test.Name)
			spew.Dump(xs, test.AlterColumns)
		}
	}
}
