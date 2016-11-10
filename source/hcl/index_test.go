package hcl

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
	"github.com/davecgh/go-spew/spew"
)

func TestConvertIndex(t *testing.T) {
	type Test struct {
		SchemaName string
		TableName  string
		IndexName  string
		Value      Index
		IndexNode  db.IndexNode
	}
	tests := []Test{
		Test{
			"public",
			"users",
			"users_email_key",
			Index{
				On: []string{"lower(email)"},
			},
			db.IndexNode{
				Index: db.Index{
					TableName: "users",
					IndexName: "users_email_key",
					Exprs:     []string{"lower(email)"},
				},
			},
		},
		Test{
			"public",
			"users",
			"users_email_key",
			Index{
				On:     []string{"email"},
				Unique: true,
			},
			db.IndexNode{
				db.Index{
					TableName: "users",
					IndexName: "users_email_key",
					Exprs:     []string{"email"},
					Unique:    true,
				},
			},
		},
	}
	for _, test := range tests {
		node := convertIndex(test.SchemaName, test.TableName, test.IndexName, test.Value)
		if !reflect.DeepEqual(node, test.IndexNode) {
			t.Errorf("bad index conversion")
			spew.Dump(node, test.IndexNode)
		}
	}
}
