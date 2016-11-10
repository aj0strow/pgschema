package psql

import (
	"fmt"
	"github.com/aj0strow/pgschema/db"
)

func LoadColumnNodes(conn Conn, schema db.Schema, table db.Table) ([]db.ColumnNode, error) {
	columns, err := LoadColumns(conn, schema.SchemaName, table.TableName)
	if err != nil {
		return nil, err
	}
	columnNodes := make([]db.ColumnNode, len(columns))
	for i := range columns {
		columnNodes[i] = db.ColumnNode{
			Column: columns[i],
		}
	}
	return columnNodes, nil
}

func LoadColumns(conn Conn, schemaName, tableName string) ([]db.Column, error) {
	q := fmt.Sprintf(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_schema = '%s'
		AND table_name = '%s'
	`, schemaName, tableName)
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := []db.Column{}
	for rows.Next() {
		column := db.Column{}
		err := rows.Scan(&column.ColumnName, &column.DataType)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return columns, nil
}
