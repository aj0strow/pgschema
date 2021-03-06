package hcl

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestConvertIndex(t *testing.T) {
	type Test struct {
		TableName string
		IndexName string
		Value     Index
		Index     *db.Index
	}
	tests := []Test{
		Test{
			"users",
			"users_email_key",
			Index{
				On: []string{"lower(email)"},
			},
			&db.Index{
				TableName: "users",
				IndexName: "users_email_key",
				Exprs:     []string{"lower(email)"},
			},
		},
		Test{
			"users",
			"users_email_key",
			Index{
				On:     []string{"email"},
				Unique: true,
			},
			&db.Index{
				TableName: "users",
				IndexName: "users_email_key",
				Exprs:     []string{"email"},
				Unique:    true,
			},
		},
	}
	for _, test := range tests {
		node := convertIndex(test.TableName, test.IndexName, test.Value)
		if !reflect.DeepEqual(node, test.Index) {
			t.Errorf("bad index conversion")
			spew.Dump(node, test.Index)
		}
	}
}
