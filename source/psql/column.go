package psql

import (
	"fmt"
	"github.com/aj0strow/pgschema/db"
)

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
