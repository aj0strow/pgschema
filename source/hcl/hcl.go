// The hcl package parses an input configuration file to create a database tree.
package hcl

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/info"
	"github.com/aj0strow/pgschema/tree"
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

func ParseBytes(bs []byte) (tree.DatabaseNode, error) {
	var db Database
	err := hcl.Unmarshal(bs, &db)
	if err != nil {
		return tree.DatabaseNode{}, err
	}
	databaseNode := convertDatabase(db)
	return databaseNode, nil
}

func convertDatabase(v Database) tree.DatabaseNode {
	var schemas []tree.SchemaNode
	for sk, sv := range v.Schema {
		schemas = append(schemas, convertSchema(sk, sv))
	}
	return tree.DatabaseNode{
		SchemaNodes: schemas,
	}
}

func convertSchema(k string, v Schema) tree.SchemaNode {
	var tables []tree.TableNode
	for tk, tv := range v.Table {
		tables = append(tables, convertTable(tk, tv))
	}
	return tree.SchemaNode{
		Schema: db.Schema{
			SchemaName: k,
		},
		TableNodes: tables,
	}
}

func convertTable(k string, v Table) tree.TableNode {
	var columns []tree.ColumnNode
	for ck, cv := range v.Column {
		columns = append(columns, convertColumn(ck, cv))
	}
	return tree.TableNode{
		Table: info.Table{
			TableName: k,
		},
		ColumnNodes: columns,
	}
}

func convertColumn(k string, v Column) tree.ColumnNode {
	return tree.ColumnNode{
		Column: info.Column{
			ColumnName: k,
			DataType:   v.Type,
		},
	}
}
