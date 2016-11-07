package pgschema

import (
	"github.com/aj0strow/pgschema/info"
	"github.com/aj0strow/pgschema/tree"
)

func LoadSchemaNode(pg PG, schema info.Schema) (tree.SchemaNode, error) {
	tables, err := info.LoadTables(pg, schema.SchemaName)
	if err != nil {
		return tree.SchemaNode{}, err
	}
	var tableNodes []tree.TableNode
	for _, table := range tables {
		tableNode, err := LoadTableNode(pg, schema, table)
		if err != nil {
			return tree.SchemaNode{}, err
		}
		tableNodes = append(tableNodes, tableNode)
	}
	schemaNode := tree.SchemaNode{
		Schema:     schema,
		TableNodes: tableNodes,
	}
	return schemaNode, nil
}

func LoadTableNode(pg PG, schema info.Schema, table info.Table) (tree.TableNode, error) {
	columns, err := info.LoadColumns(pg, schema.SchemaName, table.TableName)
	if err != nil {
		return tree.TableNode{}, err
	}
	var columnNodes []tree.ColumnNode
	for _, column := range columns {
		columnNodes = append(columnNodes, tree.ColumnNode{
			Column: column,
		})
	}
	tableNode := tree.TableNode{
		Table:       table,
		ColumnNodes: columnNodes,
	}
	return tableNode, nil
}
