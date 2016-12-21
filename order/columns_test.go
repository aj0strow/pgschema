package order

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/plan"
)

func TestAddColumnSlice(t *testing.T) {
	tests := []struct {
		Name       string
		AddColumns []plan.AddColumn
		Changes    []Change
	}{
		{
			`empty columns slice`,
			nil,
			nil,
		},
		{
			`include column type`,
			[]plan.AddColumn{
				plan.AddColumn{
					&db.Column{
						ColumnName: "email",
						DataType:   "text",
					},
				},
			},
			[]Change{
				AddColumn{
					ColumnName: "email",
					DataType:   "text",
				},
			},
		},
		{
			`include numeric precision and scale`,
			[]plan.AddColumn{
				plan.AddColumn{
					&db.Column{
						DataType:         "numeric",
						NumericPrecision: 8,
						NumericScale:     2,
					},
				},
			},
			[]Change{
				AddColumn{
					DataType:         "numeric",
					NumericPrecision: 8,
					NumericScale:     2,
				},
			},
		},
		{
			`ignore cast expression`,
			[]plan.AddColumn{
				plan.AddColumn{
					&db.Column{
						ColumnName:    "email",
						DataType:      "text",
						CastTypeUsing: "email::text",
					},
				},
			},
			[]Change{
				AddColumn{
					ColumnName: "email",
					DataType:   "text",
				},
			},
		},
		{
			`include not null constraint`,
			[]plan.AddColumn{
				plan.AddColumn{
					&db.Column{
						ColumnName: "email",
						DataType:   "text",
						NotNull:    true,
					},
				},
			},
			[]Change{
				AddColumn{
					ColumnName: "email",
					DataType:   "text",
					NotNull:    true,
				},
			},
		},
		{
			`include column default value`,
			[]plan.AddColumn{
				plan.AddColumn{
					&db.Column{
						ColumnName: "created",
						DataType:   "timestamptz",
						Default:    "now()",
					},
				},
			},
			[]Change{
				AddColumn{
					ColumnName: "created",
					DataType:   "timestamptz",
					Default:    "now()",
				},
			},
		},
	}
	for _, test := range tests {
		xs := addColumnSlice(test.AddColumns)
		if !reflect.DeepEqual(test.Changes, xs) {
			t.Errorf("addColumnSlice => %s", test.Name)
		}
	}
}

func TestAlterColumnSlice(t *testing.T) {
	tests := []struct {
		Name         string
		AlterColumns []plan.AlterColumn
		Changes      []Change
	}{
		{
			`empty columns slice`,
			nil,
			nil,
		},
		{
			`alter columns by name`,
			[]plan.AlterColumn{
				plan.AlterColumn{
					Column: &db.Column{
						ColumnName: "price",
						Default:    "0",
					},
					SetDefault: true,
				},
			},
			[]Change{
				AlterColumn{
					ColumnName: "price",
					Change:     SetDefault{"0"},
				},
			},
		},
	}
	for _, test := range tests {
		xs := alterColumnSlice(test.AlterColumns)
		if !reflect.DeepEqual(test.Changes, xs) {
			t.Errorf("alterColumnSlice => %s", test.Name)
		}
	}
}

func TestAlterColumnStruct(t *testing.T) {
	tests := []struct {
		Name        string
		AlterColumn plan.AlterColumn
		Changes     []Change
	}{
		{
			`drop default value`,
			plan.AlterColumn{
				DropDefault: true,
			},
			[]Change{
				DropDefault,
			},
		},
		{
			`drop not null constraint`,
			plan.AlterColumn{
				DropNotNull: true,
			},
			[]Change{
				DropNotNull,
			},
		},
		{
			`set data type using cast expression`,
			plan.AlterColumn{
				Column: &db.Column{
					DataType:      "text",
					CastTypeUsing: "n::text",
				},
				SetDataType: true,
			},
			[]Change{
				SetDataType{
					DataType: "text",
					Using:    "n::text",
				},
			},
		},
		{
			`set data type without expression`,
			plan.AlterColumn{
				Column: &db.Column{
					DataType: "money",
				},
				SetDataType: true,
			},
			[]Change{
				SetDataType{
					DataType: "money",
				},
			},
		},
		{
			`set data type with numeric precision and scale`,
			plan.AlterColumn{
				Column: &db.Column{
					DataType:         "numeric",
					NumericPrecision: 8,
					NumericScale:     2,
				},
				SetDataType: true,
			},
			[]Change{
				SetDataType{
					DataType:         "numeric",
					NumericPrecision: 8,
					NumericScale:     2,
				},
			},
		},
		{
			`set default value`,
			plan.AlterColumn{
				Column: &db.Column{
					Default: "0",
				},
				SetDefault: true,
			},
			[]Change{
				SetDefault{"0"},
			},
		},
		{
			`set not null`,
			plan.AlterColumn{
				SetNotNull: true,
			},
			[]Change{
				SetNotNull,
			},
		},
		{
			`set default before set not null`,
			plan.AlterColumn{
				Column: &db.Column{
					Default: "0",
				},
				SetDefault: true,
				SetNotNull: true,
			},
			[]Change{
				SetDefault{"0"},
				SetNotNull,
			},
		},
		{
			`drop default before drop not null`,
			plan.AlterColumn{
				DropDefault: true,
				DropNotNull: true,
			},
			[]Change{
				DropDefault,
				DropNotNull,
			},
		},
	}
	for _, test := range tests {
		xs := alterColumnStruct(test.AlterColumn)
		if !reflect.DeepEqual(test.Changes, xs) {
			t.Errorf("alterColumnStruct => %s", test.Name)
		}
	}
}

func TestDropColumnSlice(t *testing.T) {
	tests := []struct {
		Name        string
		DropColumns []plan.DropColumn
		Changes     []Change
	}{
		{
			`empty slice`,
			nil,
			nil,
		},
		{
			`drop columns`,
			[]plan.DropColumn{
				plan.DropColumn{
					Column: &db.Column{
						ColumnName: "color",
					},
				},
			},
			[]Change{
				DropColumn{
					ColumnName: "color",
				},
			},
		},
	}
	for _, test := range tests {
		xs := dropColumnSlice(test.DropColumns)
		if !reflect.DeepEqual(test.Changes, xs) {
			t.Errorf("dropColumnSlice => %s", test.Name)
		}
	}
}
