package pgschema

import (
	"github.com/aj0strow/pgschema/info"
)

type SchemaNode struct {
	Schema     info.Schema
	TableNodes []TableNode
}

type TableNode struct {
	Table       info.Table
	ColumnNodes []ColumnNode
}

type ColumnNode struct {
	Column info.Column
}

func LoadSchemaNode(pg PG, schema info.Schema) (SchemaNode, error) {
	tables, err := info.LoadTables(pg, schema.SchemaName)
	if err != nil {
		return SchemaNode{}, err
	}
	var tableNodes []TableNode
	for _, table := range tables {
		tableNode, err := LoadTableNode(pg, schema, table)
		if err != nil {
			return SchemaNode{}, err
		}
		tableNodes = append(tableNodes, tableNode)
	}
	schemaNode := SchemaNode{
		Schema:     schema,
		TableNodes: tableNodes,
	}
	return schemaNode, nil
}

func LoadTableNode(pg PG, schema info.Schema, table info.Table) (TableNode, error) {
	columns, err := info.LoadColumns(pg, schema.SchemaName, table.TableName)
	if err != nil {
		return TableNode{}, err
	}
	var columnNodes []ColumnNode
	for _, column := range columns {
		columnNodes = append(columnNodes, ColumnNode{
			Column: column,
		})
	}
	tableNode := TableNode{
		Table:       table,
		ColumnNodes: columnNodes,
	}
	return tableNode, nil
}
