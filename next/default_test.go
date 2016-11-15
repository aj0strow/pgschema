package next

import (
	"testing"

	"github.com/aj0strow/pgschema/db"
)

func TestSetDefault(t *testing.T) {
	type Test struct {
		Name       string
		A          *db.Column
		B          *db.Column
		SetDefault bool
	}
	tests := []Test{
		Test{
			`ignore existing defaults`,
			&db.Column{
				Default: "now()",
			},
			&db.Column{
				Default: "now()",
			},
			false,
		},
		Test{
			`ignore missing defaults`,
			&db.Column{},
			&db.Column{},
			false,
		},
		Test{
			`ignore old defaults`,
			&db.Column{},
			&db.Column{
				Default: "0",
			},
			false,
		},
		Test{
			`set new default`,
			&db.Column{
				Default: "0",
			},
			&db.Column{},
			true,
		},
	}
	for _, test := range tests {
		if setDefault(test.A, test.B) != test.SetDefault {
			t.Errorf("setDefault => %s", test.Name)
		}
	}
}

func TestDropDefault(t *testing.T) {
	type Test struct {
		Name        string
		A           *db.Column
		B           *db.Column
		DropDefault bool
	}
	tests := []Test{
		Test{
			`ignore existing defaults`,
			&db.Column{
				Default: "now()",
			},
			&db.Column{
				Default: "now()",
			},
			false,
		},
		Test{
			`ignore missing defaults`,
			&db.Column{},
			&db.Column{},
			false,
		},
		Test{
			`drop old defaults`,
			&db.Column{},
			&db.Column{
				Default: "0",
			},
			true,
		},
		Test{
			`ignore new defaults`,
			&db.Column{
				Default: "now()",
			},
			&db.Column{},
			false,
		},
	}
	for _, test := range tests {
		if dropDefault(test.A, test.B) != test.DropDefault {
			t.Errorf("dropDefault => %s", test.Name)
		}
	}
}
