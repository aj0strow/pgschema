package pgschema

import (
	"fmt"
	"github.com/jackc/pgx"
)

type PG interface {
	Query(string, ...interface{}) (*pgx.Rows, error)
	Exec(string, ...interface{}) (pgx.CommandTag, error)
}

func LoadColumns(pg PG, schemaName, tableName string) ([]Column, error) {
	q := fmt.Sprintf(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_schema = '%s'
		AND table_name = '%s'
	`, schemaName, tableName)
	rows, err := pg.Query(q)
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
