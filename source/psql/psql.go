// The psql package queries postgres meta tables to create a database db.
package psql

import (
	"github.com/aj0strow/pgschema/db"
)

func LoadDatabaseNode(conn Conn) (db.DatabaseNode, error) {
	extensionNodes, err := LoadExtensionNodes(conn)
	if err != nil {
		return db.DatabaseNode{}, err
	}
	schemas, err := LoadSchemas(conn)
	if err != nil {
		return db.DatabaseNode{}, err
	}
	var schemaNodes []db.SchemaNode
	for _, schema := range schemas {
		schemaNode, err := LoadSchemaNode(conn, schema)
		if err != nil {
			return db.DatabaseNode{}, err
		}
		schemaNodes = append(schemaNodes, schemaNode)
	}
	databaseNode := db.DatabaseNode{
		ExtensionNodes: extensionNodes,
		SchemaNodes:    schemaNodes,
	}
	return databaseNode, nil
}

func LoadSchemaNode(conn Conn, schema db.Schema) (db.SchemaNode, error) {
	tables, err := LoadTables(conn, schema.SchemaName)
	if err != nil {
		return db.SchemaNode{}, err
	}
	var tableNodes []db.TableNode
	for _, table := range tables {
		tableNode, err := LoadTableNode(conn, schema, table)
		if err != nil {
			return db.SchemaNode{}, err
		}
		tableNodes = append(tableNodes, tableNode)
	}
	schemaNode := db.SchemaNode{
		Schema:     schema,
		TableNodes: tableNodes,
	}
	return schemaNode, nil
}

func LoadTableNode(conn Conn, schema db.Schema, table db.Table) (db.TableNode, error) {
	columns, err := LoadColumns(conn, schema.SchemaName, table.TableName)
	if err != nil {
		return db.TableNode{}, err
	}
	var columnNodes []db.ColumnNode
	for _, column := range columns {
		columnNodes = append(columnNodes, db.ColumnNode{
			Column: column,
		})
	}
	tableNode := db.TableNode{
		Table:       table,
		ColumnNodes: columnNodes,
	}
	return tableNode, nil
}
