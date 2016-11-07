package info

import (
	"fmt"
)

// Table is a database table in the current search path.
type Table struct {
	TableName string
}

func LoadTables(db Conn, schemaName string) ([]Table, error) {
	q := fmt.Sprintf(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = '%s'
		AND table_type = 'BASE TABLE'
		ORDER BY table_name ASC
	`, schemaName)
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables := []Table{}
	for rows.Next() {
		table := Table{}
		rows.Scan(&table.TableName)
		tables = append(tables, table)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}
