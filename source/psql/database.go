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
	schemaNodes, err := LoadSchemaNodes(conn)
	if err != nil {
		return db.DatabaseNode{}, err
	}
	databaseNode := db.DatabaseNode{
		ExtensionNodes: extensionNodes,
		SchemaNodes:    schemaNodes,
	}
	return databaseNode, nil
}
