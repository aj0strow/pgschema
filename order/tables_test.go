package order

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/plan"
)

func TestAlterTableStruct(t *testing.T) {
	tests := []struct {
		Name       string
		Schema     *db.Schema
		AlterTable plan.AlterTable
		Changes    []Change
	}{
		{
			`nothing to change`,
			nil,
			plan.AlterTable{},
			nil,
		},
		{
			`create indexes`,
			&db.Schema{},
			plan.AlterTable{
				Table: &db.Table{},
				CreateIndexes: []plan.CreateIndex{
					plan.CreateIndex{&db.Index{}},
				},
			},
			[]Change{
				CreateIndex{},
			},
		},
		{
			`drop indexes`,
			&db.Schema{},
			plan.AlterTable{
				Table: &db.Table{},
				DropIndexes: []plan.DropIndex{
					plan.DropIndex{&db.Index{}},
				},
			},
			[]Change{
				DropIndex{},
			},
		},
		{
			`drop columns`,
			&db.Schema{},
			plan.AlterTable{
				Table: &db.Table{},
				DropColumns: []plan.DropColumn{
					plan.DropColumn{&db.Column{}},
				},
			},
			[]Change{
				AlterTable{
					Change: DropColumn{},
				},
			},
		},
		{
			`alter columns`,
			&db.Schema{},
			plan.AlterTable{
				Table: &db.Table{},
				AlterColumns: []plan.AlterColumn{
					plan.AlterColumn{
						Column:     &db.Column{},
						SetNotNull: true,
					},
				},
			},
			[]Change{
				AlterTable{
					Change: AlterColumn{
						Change: SetNotNull,
					},
				},
			},
		},
		{
			`add columns`,
			&db.Schema{},
			plan.AlterTable{
				Table: &db.Table{},
				AddColumns: []plan.AddColumn{
					plan.AddColumn{
						&db.Column{},
					},
				},
			},
			[]Change{
				AlterTable{
					Change: AddColumn{},
				},
			},
		},
		{
			`drop indexes before alter columns`,
			&db.Schema{},
			plan.AlterTable{
				Table: &db.Table{},
				DropIndexes: []plan.DropIndex{
					plan.DropIndex{
						&db.Index{
							Primary: true,
						},
					},
				},
				AlterColumns: []plan.AlterColumn{
					plan.AlterColumn{
						Column:      &db.Column{},
						DropNotNull: true,
					},
				},
			},
			[]Change{
				AlterTable{
					Change: DropConstraint{},
				},
				AlterTable{
					Change: AlterColumn{
						Change: DropNotNull,
					},
				},
			},
		},
		{
			`add columns before create indexes`,
			&db.Schema{},
			plan.AlterTable{
				Table: &db.Table{},
				AddColumns: []plan.AddColumn{
					plan.AddColumn{
						&db.Column{},
					},
				},
				CreateIndexes: []plan.CreateIndex{
					plan.CreateIndex{
						&db.Index{},
					},
				},
			},
			[]Change{
				AlterTable{
					Change: AddColumn{},
				},
				CreateIndex{},
			},
		},
	}
	for _, test := range tests {
		xs := alterTableStruct(test.Schema, test.AlterTable)
		if !reflect.DeepEqual(xs, test.Changes) {
			t.Errorf("alterTableStruct => %s", test.Name)
		}
	}
}

func TestDropTableSlice(t *testing.T) {
	tests := []struct {
		Name       string
		Schema     *db.Schema
		DropTables []plan.DropTable
		Changes    []Change
	}{
		{
			`empty table slice`,
			nil,
			nil,
			nil,
		},
		{
			`drop multiple tables`,
			&db.Schema{"public"},
			[]plan.DropTable{
				plan.DropTable{
					Table: &db.Table{"temp1"},
				},
				plan.DropTable{
					Table: &db.Table{"temp2"},
				},
			},
			[]Change{
				DropTable{"public", "temp1"},
				DropTable{"public", "temp2"},
			},
		},
	}
	for _, test := range tests {
		xs := dropTableSlice(test.Schema, test.DropTables)
		if !reflect.DeepEqual(xs, test.Changes) {
			t.Errorf("dropTableSlice => %s", test.Name)
		}
	}
}
