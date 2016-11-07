package info

import (
	"fmt"
)

type Schema struct {
	SchemaName string
}

func LoadSchemas(db Conn) ([]Schema, error) {
	q := fmt.Sprintf(`
		SELECT schema_name
		FROM information_schema.schemata
		WHERE schema_name !~ 'pg_'
		AND schema_name <> 'information_schema'
	`)
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var schemas []Schema
	for rows.Next() {
		schema := Schema{}
		rows.Scan(&schema.SchemaName)
		schemas = append(schemas, schema)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return schemas, nil
}
