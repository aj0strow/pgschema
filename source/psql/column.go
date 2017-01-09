package psql

import (
	"fmt"

	"github.com/aj0strow/pgschema/db"
	"github.com/jackc/pgx"
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
		SELECT
			column_name,
			data_type,
			is_nullable,
			column_default,
			numeric_precision,
			numeric_scale,
			numeric_precision_radix,
			udt_name
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
		var (
			isNullable       string
			colDefault       pgx.NullString
			numericPrecision pgx.NullInt32
			numericScale     pgx.NullInt32
			numericRadix     pgx.NullInt32
			udtName          string
		)
		err := rows.Scan(
			&column.ColumnName,
			&column.DataType,
			&isNullable,
			&colDefault,
			&numericPrecision,
			&numericScale,
			&numericRadix,
			&udtName,
		)
		if err != nil {
			return nil, err
		}
		if column.DataType == "USER-DEFINED" {
			column.DataType = udtName
		}
		column.NotNull = isNullable == "NO"
		column.Default = colDefault.String
		if numericRadix.Valid {
			column.NumericPrecision = int(numericPrecision.Int32)
			column.NumericScale = int(numericScale.Int32)
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return columns, nil
}
