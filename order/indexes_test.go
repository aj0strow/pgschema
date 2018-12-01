package order

import (
	"reflect"
	"testing"

	"github.com/aj0strow/pgschema/db"
)

func TestCreateIndexStruct(t *testing.T) {
	tests := []struct {
		Name   string
		Schema *db.Schema
		Table  *db.Table
		Index  *db.Index
		Change Change
	}{
		{
			`add primary key`,
			&db.Schema{SchemaName: "public"},
			&db.Table{TableName: "users"},
			&db.Index{
				Primary: true,
				Exprs:   []string{"id"},
			},
			AlterTable{
				"public",
				"users",
				AddPrimaryKey{
					Columns: []string{"id"},
				},
			},
		},
		{
			`create unique index on expression`,
			&db.Schema{SchemaName: "public"},
			&db.Table{TableName: "users"},
			&db.Index{
				IndexName: "users_email_key",
				Exprs:     []string{"lower(email)"},
				Unique:    true,
			},
			CreateIndex{
				SchemaName: "public",
				TableName:  "users",
				IndexName:  "users_email_key",
				Exprs:      []string{"lower(email)"},
				Unique:     true,
			},
		},
		{
			`create compound index with sort order`,
			&db.Schema{SchemaName: "public"},
			&db.Table{TableName: "users"},
			&db.Index{
				IndexName: "users_active_by_city",
				Exprs:     []string{"city", "last_active_at DESC"},
			},
			CreateIndex{
				SchemaName: "public",
				TableName:  "users",
				IndexName:  "users_active_by_city",
				Exprs:      []string{"city", "last_active_at DESC"},
			},
		},
	}
	for _, test := range tests {
		x := createIndexStruct(test.Schema, test.Table, test.Index)
		if !reflect.DeepEqual(x, test.Change) {
			t.Errorf("createIndexStruct => %s", test.Name)
		}
	}
}

func TestDropIndexStruct(t *testing.T) {
	tests := []struct {
		Name   string
		Schema *db.Schema
		Table  *db.Table
		Index  *db.Index
		Change Change
	}{
		{
			`drop primary key`,
			&db.Schema{SchemaName: "public"},
			&db.Table{TableName: "users"},
			&db.Index{
				IndexName: "users_pkey",
				Primary:   true,
			},
			AlterTable{
				"public",
				"users",
				DropConstraint{"users_pkey"},
			},
		},
		{
			`drop normal index`,
			&db.Schema{SchemaName: "public"},
			&db.Table{},
			&db.Index{
				IndexName: "users_email_key",
				Unique:    true,
			},
			DropIndex{
				"public",
				"users_email_key",
			},
		},
	}
	for _, test := range tests {
		x := dropIndexStruct(test.Schema, test.Table, test.Index)
		if !reflect.DeepEqual(x, test.Change) {
			t.Errorf("dropIndexStruct => %s", test.Name)
		}
	}
}
