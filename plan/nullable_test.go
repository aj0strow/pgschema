package plan

import (
	"testing"

	"github.com/aj0strow/pgschema/db"
)

func TestSetNotNull(t *testing.T) {
	type Test struct {
		Name       string
		A          *db.Column
		B          *db.Column
		SetNotNull bool
	}
	tests := []Test{
		Test{
			`ignore existing not null`,
			&db.Column{
				NotNull: true,
			},
			&db.Column{
				NotNull: true,
			},
			false,
		},
		Test{
			`ignore existing null ok`,
			&db.Column{},
			&db.Column{},
			false,
		},
		Test{
			`ignore prior not null`,
			&db.Column{},
			&db.Column{
				NotNull: true,
			},
			false,
		},
		Test{
			`set not null when new`,
			&db.Column{
				NotNull: true,
			},
			&db.Column{},
			true,
		},
	}
	for _, test := range tests {
		if setNotNull(test.A, test.B) != test.SetNotNull {
			t.Errorf("setNotNull => %s", test.Name)
		}
	}
}

func TestDropNotNull(t *testing.T) {
	type Test struct {
		Name        string
		A           *db.Column
		B           *db.Column
		DropNotNull bool
	}
	tests := []Test{
		Test{
			`ignore existing not null`,
			&db.Column{
				NotNull: true,
			},
			&db.Column{
				NotNull: true,
			},
			false,
		},
		Test{
			`ignore existing null ok`,
			&db.Column{},
			&db.Column{},
			false,
		},
		Test{
			`drop old not null constraint`,
			&db.Column{},
			&db.Column{
				NotNull: true,
			},
			true,
		},
		Test{
			`ignore new not null constraint`,
			&db.Column{
				NotNull: true,
			},
			&db.Column{},
			false,
		},
	}
	for _, test := range tests {
		if dropNotNull(test.A, test.B) != test.DropNotNull {
			t.Errorf("dropNotNull => %s", test.Name)
		}
	}
}
