package plan

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestCreateExtensions(t *testing.T) {
	type Test struct {
		Name             string
		ExtensionMatches []ab.ExtensionMatch
		CreateExtensions []CreateExtension
	}
	tests := []Test{
		Test{
			`empty extensions list`,
			nil,
			nil,
		},
		Test{
			`add new extension`,
			[]ab.ExtensionMatch{
				ab.ExtensionMatch{
					A: &db.Extension{
						ExtName: "hstore",
					},
				},
			},
			[]CreateExtension{
				CreateExtension{
					&db.Extension{
						ExtName: "hstore",
					},
				},
			},
		},
		Test{
			`ignore existing extension`,
			[]ab.ExtensionMatch{
				ab.ExtensionMatch{
					A: &db.Extension{
						ExtName: "hstore",
					},
					B: &db.Extension{
						ExtName: "hstore",
					},
				},
			},
			nil,
		},
	}
	for _, test := range tests {
		xs := createExtensions(test.ExtensionMatches)
		if !reflect.DeepEqual(xs, test.CreateExtensions) {
			t.Errorf("createExtensions => %s", test.Name)
			spew.Dump(xs, test.CreateExtensions)
		}
	}
}

func TestDropExtensions(t *testing.T) {
	type Test struct {
		Name             string
		ExtensionMatches []ab.ExtensionMatch
		DropExtensions   []DropExtension
	}
	tests := []Test{
		Test{
			`empty extension list`,
			nil,
			nil,
		},
		Test{
			`ignore plpgsql`,
			[]ab.ExtensionMatch{
				ab.ExtensionMatch{
					B: &db.Extension{
						ExtName: "plpgsql",
					},
				},
			},
			nil,
		},
		Test{
			`drop existing extension`,
			[]ab.ExtensionMatch{
				ab.ExtensionMatch{
					B: &db.Extension{
						ExtName: "hstore",
					},
				},
			},
			[]DropExtension{
				DropExtension{
					&db.Extension{
						ExtName: "hstore",
					},
				},
			},
		},
		Test{
			`ignore required extension`,
			[]ab.ExtensionMatch{
				ab.ExtensionMatch{
					A: &db.Extension{
						ExtName: "hstore",
					},
					B: &db.Extension{
						ExtName: "hstore",
					},
				},
			},
			nil,
		},
	}
	for _, test := range tests {
		xs := dropExtensions(test.ExtensionMatches)
		if !reflect.DeepEqual(xs, test.DropExtensions) {
			t.Errorf("dropExtensions => %s", test.Name)
			spew.Dump(xs, test.DropExtensions)
		}
	}
}
