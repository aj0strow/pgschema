// The hcl package parses an input configuration file to create a database db.
package hcl

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/hashicorp/hcl"
)

type Database struct {
	Extension map[string]Extension
	Schema    map[string]Schema
}

type Extension struct{}

type Schema struct {
	Table map[string]Table
}

type Table struct {
	Column map[string]Column
	Index  map[string]Index
}

type Column struct {
	Type          string
	NotNull       bool   `hcl:"not_null"`
	CastTypeUsing string `hcl:"cast_type_using"`
}

type Index struct{}

func ParseBytes(bs []byte) (db.DatabaseNode, error) {
	var database Database
	err := hcl.Unmarshal(bs, &database)
	if err != nil {
		return db.DatabaseNode{}, err
	}
	databaseNode := convertDatabase(database)
	return databaseNode, nil
}

func convertDatabase(v Database) db.DatabaseNode {
	var exts []db.ExtensionNode
	for ek := range v.Extension {
		exts = append(exts, convertExtension(ek))
	}
	var schemas []db.SchemaNode
	for sk, sv := range v.Schema {
		schemas = append(schemas, convertSchema(sk, sv))
	}
	return db.DatabaseNode{
		ExtensionNodes: exts,
		SchemaNodes:    schemas,
	}
}

func convertExtension(k string) db.ExtensionNode {
	return db.ExtensionNode{
		Extension: db.Extension{
			ExtName: k,
		},
	}
}

func convertSchema(k string, v Schema) db.SchemaNode {
	var tables []db.TableNode
	for tk, tv := range v.Table {
		tables = append(tables, convertTable(tk, tv))
	}
	return db.SchemaNode{
		Schema: db.Schema{
			SchemaName: k,
		},
		TableNodes: tables,
	}
}

func convertTable(k string, v Table) db.TableNode {
	var columns []db.ColumnNode
	for ck, cv := range v.Column {
		columns = append(columns, convertColumn(ck, cv))
	}
	var indexes []db.IndexNode
	for ik, iv := range v.Index {
		indexes = append(indexes, convertIndex(ik, iv))
	}
	return db.TableNode{
		Table: db.Table{
			TableName: k,
		},
		ColumnNodes: columns,
		IndexNodes:  indexes,
	}
}

func convertColumn(k string, v Column) db.ColumnNode {
	return db.ColumnNode{
		Column: db.Column{
			ColumnName:    k,
			DataType:      v.Type,
			CastTypeUsing: v.CastTypeUsing,
			NotNull:       v.NotNull,
		},
	}
}

func convertIndex(k string, v Index) db.IndexNode {
	return db.IndexNode{
		Index: db.Index{
			IndexName: k,
		},
	}
}
