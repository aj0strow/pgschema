// The hcl package parses an input configuration file to create a database db.
package hcl

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/hashicorp/hcl"
)

type Database struct {
	Schema map[string]Schema
}

type Schema struct {
	Table map[string]Table
}

type Table struct {
	Column map[string]Column
}

type Column struct {
	Type string
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
	var schemas []db.SchemaNode
	for sk, sv := range v.Schema {
		schemas = append(schemas, convertSchema(sk, sv))
	}
	return db.DatabaseNode{
		SchemaNodes: schemas,
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
	return db.TableNode{
		Table: db.Table{
			TableName: k,
		},
		ColumnNodes: columns,
	}
}

func convertColumn(k string, v Column) db.ColumnNode {
	return db.ColumnNode{
		Column: db.Column{
			ColumnName: k,
			DataType:   v.Type,
		},
	}
}
