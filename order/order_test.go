package order

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/next"
)

func TestChanges(t *testing.T) {
	type Test struct {
		Name           string
		UpdateDatabase next.UpdateDatabase
		Changes        []Change
	}
	tests := []Test{
		Test{
			`drop tables`,
			next.UpdateDatabase{
				UpdateSchemas: []next.UpdateSchema{
					next.UpdateSchema{
						Schema: &db.Schema{
							SchemaName: "public",
						},
						DropTables: []next.DropTable{
							next.DropTable{
								Table: &db.Table{
									TableName: "users",
								},
							},
						},
					},
				},
			},
			[]Change{
				DropTable{
					SchemaName: "public",
					TableName:  "users",
				},
			},
		},
		Test{
			`drop primary key constraints`,
			next.UpdateDatabase{
				UpdateSchemas: []next.UpdateSchema{
					next.UpdateSchema{
						Schema: &db.Schema{
							SchemaName: "public",
						},
						AlterTables: []next.AlterTable{
							next.AlterTable{
								Table: &db.Table{
									TableName: "users",
								},
								DropIndexes: []next.DropIndex{
									next.DropIndex{
										Index: &db.Index{
											IndexName: "users_pkey",
											Primary:   true,
										},
									},
								},
							},
						},
					},
				},
			},
			[]Change{
				AlterTable{
					"public",
					"users",
					DropConstraint{"users_pkey"},
				},
			},
		},
	}
	for _, test := range tests {
		xs := Changes(test.UpdateDatabase)
		if !reflect.DeepEqual(xs, test.Changes) {
			t.Errorf("Changes => %s", test.Name)
		}
	}
}
