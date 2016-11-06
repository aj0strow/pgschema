package pgschema

import (
	"fmt"
	"github.com/jackc/pgx"
)

type PG interface {
	Query(string, ...interface{}) (*pgx.Rows, error)
	Exec(string, ...interface{}) (pgx.CommandTag, error)
}

func LoadTables(pg PG, schemaName string) ([]Table, error) {
	q := fmt.Sprintf(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = '%s'
		AND table_type = 'BASE TABLE'
		ORDER BY table_name ASC
	`, schemaName)
	rows, err := pg.Query(q)
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
