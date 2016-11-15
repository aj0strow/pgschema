package next

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestCreateSchemas(t *testing.T) {
	type Test struct {
		Name          string
		SchemaMatches []ab.SchemaMatch
		CreateSchemas []CreateSchema
	}
	tests := []Test{
		Test{
			`empty schema list`,
			nil,
			nil,
		},
		Test{
			`create required schema`,
			[]ab.SchemaMatch{
				ab.SchemaMatch{
					A: &db.Schema{
						SchemaName: "v1",
					},
				},
			},
			[]CreateSchema{
				CreateSchema{
					Schema: &db.Schema{
						SchemaName: "v1",
					},
				},
			},
		},
		Test{
			`ignore existing schema`,
			[]ab.SchemaMatch{
				ab.SchemaMatch{
					A: &db.Schema{
						SchemaName: "v1",
					},
					B: &db.Schema{
						SchemaName: "v1",
					},
				},
			},
			nil,
		},
	}
	for _, test := range tests {
		xs := createSchemas(test.SchemaMatches)
		if !reflect.DeepEqual(xs, test.CreateSchemas) {
			t.Errorf("createSchemas => %s", test.Name)
			spew.Dump(xs, test.CreateSchemas)
		}
	}
}

func TestUpdateSchemas(t *testing.T) {
	type Test struct {
		Name          string
		SchemaMatches []ab.SchemaMatch
		UpdateSchemas []UpdateSchema
	}
	tests := []Test{
		Test{
			`empty schema list`,
			nil,
			nil,
		},
		Test{
			`ignore mismatched schemas`,
			[]ab.SchemaMatch{
				ab.SchemaMatch{
					A: &db.Schema{},
				},
			},
			nil,
		},
	}
	for _, test := range tests {
		xs := updateSchemas(test.SchemaMatches)
		if !reflect.DeepEqual(xs, test.UpdateSchemas) {
			t.Errorf("updateSchemas => %s", test.Name)
			spew.Dump(xs, test.UpdateSchemas)
		}
	}
}
