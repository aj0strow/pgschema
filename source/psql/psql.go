// The psql package queries postgres meta tables to create a database db.
package psql

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/info"
)

func LoadDatabaseNode(conn info.Conn) (db.DatabaseNode, error) {
	schemas, err := info.LoadSchemas(conn)
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
		SchemaNodes: schemaNodes,
	}
	return databaseNode, nil
}

func LoadSchemaNode(conn info.Conn, schema db.Schema) (db.SchemaNode, error) {
	tables, err := info.LoadTables(conn, schema.SchemaName)
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

func LoadTableNode(conn info.Conn, schema db.Schema, table db.Table) (db.TableNode, error) {
	columns, err := info.LoadColumns(conn, schema.SchemaName, table.TableName)
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
