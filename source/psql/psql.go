// The psql package queries postgres meta tables to create a database tree.
package psql

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/info"
	"github.com/aj0strow/pgschema/tree"
)

func LoadDatabaseNode(db info.Conn) (tree.DatabaseNode, error) {
	schemas, err := info.LoadSchemas(db)
	if err != nil {
		return tree.DatabaseNode{}, err
	}
	var schemaNodes []tree.SchemaNode
	for _, schema := range schemas {
		schemaNode, err := LoadSchemaNode(db, schema)
		if err != nil {
			return tree.DatabaseNode{}, err
		}
		schemaNodes = append(schemaNodes, schemaNode)
	}
	databaseNode := tree.DatabaseNode{
		SchemaNodes: schemaNodes,
	}
	return databaseNode, nil
}

func LoadSchemaNode(db info.Conn, schema db.Schema) (tree.SchemaNode, error) {
	tables, err := info.LoadTables(db, schema.SchemaName)
	if err != nil {
		return tree.SchemaNode{}, err
	}
	var tableNodes []tree.TableNode
	for _, table := range tables {
		tableNode, err := LoadTableNode(db, schema, table)
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

func LoadTableNode(db info.Conn, schema db.Schema, table info.Table) (tree.TableNode, error) {
	columns, err := info.LoadColumns(db, schema.SchemaName, table.TableName)
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
