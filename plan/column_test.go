package plan

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestColumnChanges(t *testing.T) {
	type Test struct {
		Name        string
		ColumnMatch ab.ColumnMatch
		Changes     []Change
	}
	tests := []Test{
		Test{
			"drop old column",
			ab.ColumnMatch{
				A: nil,
				B: &db.Column{
					ColumnName: "email",
				},
			},
			[]Change{
				DropColumn{"email"},
			},
		},
		Test{
			"add new column",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName: "email",
					DataType:   "citext",
				},
				B: nil,
			},
			[]Change{
				AddColumn{"email", "citext"},
			},
		},
		Test{
			"change column type",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName: "email",
					DataType:   "text",
				},
				B: &db.Column{
					ColumnName: "email",
					DataType:   "citext",
				},
			},
			[]Change{
				AlterColumn{
					"email",
					SetDataType{"text"},
				},
			},
		},
		Test{
			"cast column type",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName:    "cost",
					DataType:      "money",
					CastTypeUsing: "(cost * 100)::money",
				},
				B: &db.Column{
					ColumnName: "cost",
					DataType:   "integer",
				},
			},
			[]Change{
				AlterColumn{
					"cost",
					CastDataType{
						SetDataType{"money"},
						"(cost * 100)::money",
					},
				},
			},
		},
		Test{
			"set not null",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName: "cost",
					NotNull:    true,
				},
				B: &db.Column{
					ColumnName: "cost",
				},
			},
			[]Change{
				AlterColumn{
					"cost",
					SetNotNull{},
				},
			},
		},
		Test{
			"drop not null",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName: "cost",
				},
				B: &db.Column{
					ColumnName: "cost",
					NotNull:    true,
				},
			},
			[]Change{
				AlterColumn{
					"cost",
					DropNotNull{},
				},
			},
		},
		Test{
			"set default",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName: "cost",
					Default:    "0",
				},
				B: &db.Column{
					ColumnName: "cost",
				},
			},
			[]Change{
				AlterColumn{
					"cost",
					SetDefault{"0"},
				},
			},
		},
		Test{
			"drop default",
			ab.ColumnMatch{
				A: &db.Column{
					ColumnName: "cost",
				},
				B: &db.Column{
					ColumnName: "cost",
					Default:    "0",
				},
			},
			[]Change{
				AlterColumn{
					"cost",
					DropDefault{},
				},
			},
		},
	}
	for _, test := range tests {
		changes := ColumnChanges(test.ColumnMatch)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("planColumnMatch => %s", test.Name)
			spew.Dump(changes, test.Changes)
		}
	}
}
