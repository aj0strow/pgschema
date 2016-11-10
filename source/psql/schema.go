package psql

import (
	"fmt"
	"github.com/aj0strow/pgschema/db"
)

func LoadSchemaNodes(conn Conn) ([]db.SchemaNode, error) {
	schemas, err := LoadSchemas(conn)
	if err != nil {
		return nil, err
	}
	schemaNodes := make([]db.SchemaNode, len(schemas))
	for i := range schemas {
		tableNodes, err := LoadTableNodes(conn, schemas[i])
		if err != nil {
			return nil, err
		}
		schemaNodes[i] = db.SchemaNode{
			Schema:     schemas[i],
			TableNodes: tableNodes,
		}
	}
	return schemaNodes, nil
}

func LoadSchemas(conn Conn) ([]db.Schema, error) {
	q := fmt.Sprintf(`
		SELECT schema_name
		FROM information_schema.schemata
		WHERE schema_name !~ 'pg_'
		AND schema_name <> 'information_schema'
	`)
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var schemas []db.Schema
	for rows.Next() {
		schema := db.Schema{}
		rows.Scan(&schema.SchemaName)
		schemas = append(schemas, schema)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return schemas, nil
}
