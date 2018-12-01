package psql

import (
	"fmt"

	"github.com/aj0strow/pgschema/db"
	"github.com/jackc/pgx/pgtype"
)

func LoadColumns(conn Conn, schemaName, tableName string) ([]*db.Column, error) {
	elementTypes, err := LoadElementTypes(conn, schemaName, tableName)
	if err != nil {
		return nil, err
	}
	q := fmt.Sprintf(`
		SELECT
			column_name,
			data_type,
			is_nullable,
			column_default,
			numeric_precision,
			numeric_scale,
			numeric_precision_radix,
			udt_name,
			dtd_identifier
		FROM information_schema.columns
		WHERE table_schema = '%s'
		AND table_name = '%s'
	`, schemaName, tableName)
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := []*db.Column{}
	for rows.Next() {
		column := &db.Column{}
		var (
			isNullable       string
			colDefault       pgtype.Text
			numericPrecision pgtype.Int4
			numericScale     pgtype.Int4
			numericRadix     pgtype.Int4
			udtName          string
			elementTypeId    string
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
			&elementTypeId,
		)
		if err != nil {
			return nil, err
		}
		column.NotNull = isNullable == "NO"
		column.Default = colDefault.String
		if numericRadix.Status == pgtype.Present {
			column.NumericPrecision = int(numericPrecision.Int)
			column.NumericScale = int(numericScale.Int)
		}
		if column.DataType == "USER-DEFINED" {
			column.DataType = udtName
		}
		if column.DataType == "ARRAY" {
			elementType := findElementType(elementTypes, elementTypeId)
			column.Array = true
			if elementType != nil {
				column.DataType = elementType.DataType
				column.NumericPrecision = elementType.NumericPrecision
				column.NumericScale = elementType.NumericScale
			}
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return columns, nil
}

func findElementType(elementTypes []*ElementType, elementId string) *ElementType {
	for i := range elementTypes {
		if elementTypes[i].ElementTypeId == elementId {
			return elementTypes[i]
		}
	}
	return nil
}

type ElementType struct {
	ElementTypeId    string
	DataType         string
	NumericPrecision int
	NumericScale     int
}

func LoadElementTypes(conn Conn, schemaName, tableName string) ([]*ElementType, error) {
	q := fmt.Sprintf(`
		SELECT
			collection_type_identifier,
			data_type,
			numeric_precision,
			numeric_scale,
			numeric_precision_radix,
			udt_name
		FROM information_schema.element_types
		WHERE object_schema = '%s'
		AND object_name = '%s'
		AND object_type = 'TABLE'
	`, schemaName, tableName)
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	elementTypes := []*ElementType{}
	for rows.Next() {
		elementType := &ElementType{}
		var (
			numericPrecision pgtype.Int4
			numericScale     pgtype.Int4
			numericRadix     pgtype.Int4
			udtName          string
		)
		err := rows.Scan(
			&elementType.ElementTypeId,
			&elementType.DataType,
			&numericPrecision,
			&numericScale,
			&numericRadix,
			&udtName,
		)
		if err != nil {
			return nil, err
		}
		if elementType.DataType == "USER-DEFINED" {
			elementType.DataType = udtName
		}
		if numericRadix.Status == pgtype.Present {
			elementType.NumericPrecision = int(numericPrecision.Int)
			elementType.NumericScale = int(numericScale.Int)
		}
		elementTypes = append(elementTypes, elementType)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return elementTypes, nil
}
