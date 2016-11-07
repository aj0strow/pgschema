package info

import (
	"fmt"
)

type Column struct {
	ColumnName string
	DataType   string
}

func LoadColumns(db Conn, schemaName, tableName string) ([]Column, error) {
	q := fmt.Sprintf(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_schema = '%s'
		AND table_name = '%s'
	`, schemaName, tableName)
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := []Column{}
	for rows.Next() {
		column := Column{}
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
