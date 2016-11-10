package plan

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

func TestIndexChanges(t *testing.T) {
	type Test struct {
		Name       string
		IndexMatch ab.IndexMatch
		Changes    []Change
	}
	tests := []Test{
		Test{
			"create new index",
			ab.IndexMatch{
				A: &db.Index{
					TableName: "users",
					IndexName: "users_email_key",
					Exprs:     []string{"lower(email)"},
				},
			},
			[]Change{
				CreateIndex{
					TableName: "users",
					IndexName: "users_email_key",
					Exprs:     []string{"lower(email)"},
				},
			},
		},
		Test{
			"drop existing index",
			ab.IndexMatch{
				B: &db.Index{
					IndexName: "users_pkey",
				},
			},
			[]Change{
				DropIndex{"users_pkey"},
			},
		},
		Test{
			"existing index noop",
			ab.IndexMatch{
				A: &db.Index{
					IndexName: "users_email_key",
					Exprs:     []string{"lower(email)"},
				},
				B: &db.Index{
					IndexName: "users_email_key",
				},
			},
			nil,
		},
		Test{
			"change index unique",
			ab.IndexMatch{
				A: &db.Index{
					TableName: "users",
					IndexName: "users_email_key",
					Exprs:     []string{"lower(email)"},
					Unique:    true,
				},
				B: &db.Index{
					TableName: "users",
					IndexName: "users_email_key",
					Exprs:     []string{"lower(email)"},
				},
			},
			[]Change{
				DropIndex{"users_email_key"},
				CreateIndex{
					TableName: "users",
					IndexName: "users_email_key",
					Exprs:     []string{"lower(email)"},
					Unique:    true,
				},
			},
		},
		Test{
			"add primary key",
			ab.IndexMatch{
				A: &db.Index{
					TableName: "users",
					IndexName: "users_pkey",
					Exprs:     []string{"email"},
					Unique:    true,
					Primary:   true,
				},
			},
			[]Change{
				AlterTable{
					"users",
					AddPrimaryKey{"users_pkey", []string{"email"}},
				},
			},
		},
		Test{
			"drop primary key",
			ab.IndexMatch{
				B: &db.Index{
					TableName: "users",
					IndexName: "users_pkey",
					Exprs:     []string{"email"},
					Unique:    true,
					Primary:   true,
				},
			},
			[]Change{
				AlterTable{
					"users",
					DropConstraint{"users_pkey"},
				},
			},
		},
	}
	for _, test := range tests {
		changes := IndexChanges(test.IndexMatch)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("IndexChanges => %s", test.Name)
		}
	}
}
