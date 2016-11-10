// The hcl package parses an input configuration file to create a database db.
package hcl

import (
	"fmt"

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
	PrimaryKey []string `hcl:"primary_key"`
	Column     map[string]Column
	Index      map[string]Index
}

type Column struct {
	Type          string
	NotNull       bool   `hcl:"not_null"`
	CastTypeUsing string `hcl:"cast_type_using"`
	Default       string
	PrimaryKey    bool `hcl:"primary_key"`
}

type Index struct {
	On     []string
	Unique bool
}

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
	for tableName, tv := range v.Table {
		tables = append(tables, convertTable(tableName, tv))
	}
	return db.SchemaNode{
		Schema: db.Schema{
			SchemaName: k,
		},
		TableNodes: tables,
	}
}

func convertTable(tableName string, v Table) db.TableNode {
	var (
		columns []db.ColumnNode
		indexes []db.IndexNode
	)
	for columnName, c := range v.Column {
		columns = append(columns, convertColumn(columnName, c))
		if c.PrimaryKey {
			indexes = append(indexes, newPrimaryKey(tableName, []string{columnName}))
		}
	}
	if len(v.PrimaryKey) > 0 {
		indexes = append(indexes, newPrimaryKey(tableName, v.PrimaryKey))
	}
	for indexName, ix := range v.Index {
		indexes = append(indexes, convertIndex(tableName, indexName, ix))
	}
	return db.TableNode{
		Table: db.Table{
			TableName: tableName,
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
			NotNull:       v.NotNull || v.PrimaryKey,
			Default:       v.Default,
		},
	}
}

func convertIndex(tableName, indexName string, v Index) db.IndexNode {
	return db.IndexNode{
		Index: db.Index{
			TableName: tableName,
			IndexName: indexName,
			Exprs:     v.On,
			Unique:    v.Unique,
		},
	}
}

func newPrimaryKey(tableName string, columnNames []string) db.IndexNode {
	return db.IndexNode{
		Index: db.Index{
			TableName: tableName,
			IndexName: fmt.Sprintf("%s_pkey", tableName),
			Exprs:     columnNames,
			Unique:    true,
			Primary:   true,
		},
	}
}
