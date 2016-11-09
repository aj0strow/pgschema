package info

import (
	"fmt"
	"github.com/aj0strow/pgschema/db"
)

func LoadTables(conn Conn, schemaName string) ([]db.Table, error) {
	q := fmt.Sprintf(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = '%s'
		AND table_type = 'BASE TABLE'
		ORDER BY table_name ASC
	`, schemaName)
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables := []db.Table{}
	for rows.Next() {
		table := db.Table{}
		rows.Scan(&table.TableName)
		tables = append(tables, table)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}
