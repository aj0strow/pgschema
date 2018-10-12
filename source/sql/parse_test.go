package sql

import (
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	cs := []struct {
		Name         string
		SQL          string
		DatabaseNode db.DatabaseNode
	}{
		{
			`empty file`,
			``,
			db.DatabaseNode{},
		},
		{
			"parse extensions",
			`EXTENSION 'citext';`,
			db.DatabaseNode{
				ExtensionNodes: []db.ExtensionNode{
					db.ExtensionNode{
						Extension: db.Extension{
							ExtName: "citext",
						},
					},
				},
			},
		},
		{
			"parse schema declaration",
			"SCHEMA default;",
			db.DatabaseNode{
				SchemaNodes: []db.SchemaNode{
					db.SchemaNode{
						Schema: db.Schema{
							SchemaName: "default",
						},
					},
				},
			},
		},
	}
	for _, c := range cs {
		n, err := parse(c.SQL)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(n, c.DatabaseNode) {
			t.Errorf(c.Name)
		}
	}
}
