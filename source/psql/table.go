package psql

import (
	"fmt"
	"github.com/aj0strow/pgschema/db"
)

func LoadTableNodes(conn Conn, schema db.Schema) ([]db.TableNode, error) {
	tables, err := LoadTables(conn, schema.SchemaName)
	if err != nil {
		return nil, err
	}
	tableNodes := make([]db.TableNode, len(tables))
	for i := range tables {
		columnNodes, err := LoadColumnNodes(conn, schema, tables[i])
		if err != nil {
			return nil, err
		}
		indexNodes, err := LoadIndexNodes(conn, schema, tables[i])
		tableNodes[i] = db.TableNode{
			Table:       tables[i],
			ColumnNodes: columnNodes,
			IndexNodes:  indexNodes,
		}
	}
	return tableNodes, nil
}

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
