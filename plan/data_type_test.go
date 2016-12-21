package plan

import (
	"testing"

	"github.com/aj0strow/pgschema/db"
)

func TestSetDataType(t *testing.T) {
	type Test struct {
		Name        string
		A           *db.Column
		B           *db.Column
		SetDataType bool
	}
	tests := []Test{
		Test{
			`ignore same type`,
			&db.Column{
				DataType: "text",
			},
			&db.Column{
				DataType: "text",
			},
			false,
		},
		Test{
			`change different type`,
			&db.Column{
				ColumnName:    "created_at",
				DataType:      "timestamptz",
				CastTypeUsing: "created_at AT TIME ZONE 'UTC'",
			},
			&db.Column{
				DataType: "timestamp",
			},
			true,
		},
		Test{
			`change numeric to double`,
			&db.Column{
				DataType: "numeric",
			},
			&db.Column{
				DataType: "double precision",
			},
			true,
		},
		Test{
			`ignore timestamp alias`,
			&db.Column{
				DataType: "timestamp",
			},
			&db.Column{
				DataType: "timestamp without time zone",
			},
			false,
		},
		Test{
			`ignore timestamptz alias`,
			&db.Column{
				DataType: "timestamptz",
			},
			&db.Column{
				DataType: "timestamp with time zone",
			},
			false,
		},
	}
	for _, test := range tests {
		if setDataType(test.A, test.B) != test.SetDataType {
			t.Errorf("setDataType => %s", test.Name)
		}
	}
}
