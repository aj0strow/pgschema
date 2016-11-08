package plan

import (
	"fmt"
)

// Change is a single atomic schema change.
type Change interface {
	String() string
}

type CreateSchema struct {
	SchemaName string
}

func (cs CreateSchema) String() string {
	return fmt.Sprintf(`CREATE SCHEMA %s`, cs.SchemaName)
}

var _ Change = (*CreateSchema)(nil)

// Create a new table. This change occurs when you have a table name in
// the new schema with no match in the old schema.
type CreateTable struct {
	TableName string
}

func (ct CreateTable) String() string {
	return fmt.Sprintf(`CREATE TABLE %s ()`, ct.TableName)
}

var _ Change = (*CreateTable)(nil)

// Drop an existing table. This change occurs when you have a table name
// in the old schema with no match in the new schema.
type DropTable struct {
	TableName string
}

func (dt DropTable) String() string {
	return fmt.Sprintf(`DROP TABLE %s`, dt.TableName)
}

var _ Change = (*DropTable)(nil)

// Alter an existing table. This change occurs when you have matching
// table names but the schema doesn't match.
type AlterTable struct {
	TableName string
	Change    Change
}

func (at AlterTable) String() string {
	return fmt.Sprintf(`ALTER TABLE %s %s`, at.TableName, at.Change)
}

var _ Change = (*AlterColumn)(nil)

// Add a new column to an existing table.
type AddColumn struct {
	ColumnName string
	DataType   string
}

func (ac AddColumn) String() string {
	return fmt.Sprintf(`ADD COLUMN %s %s`, ac.ColumnName, ac.DataType)
}

var _ Change = (*AddColumn)(nil)

// Drop existing column from existing table.
type DropColumn struct {
	ColumnName string
}

func (dc DropColumn) String() string {
	return fmt.Sprintf(`DROP COLUMN %s RESTRICT`, dc.ColumnName)
}

var _ Change = (*DropColumn)(nil)

// Atler an existing column.
type AlterColumn struct {
	ColumnName string
	Change     Change
}

func (ac AlterColumn) String() string {
	return fmt.Sprintf(`ALTER COLUMN %s %s`, ac.ColumnName, ac.Change)
}

var _ Change = (*AlterColumn)(nil)

// Change the column data type.
type SetDataType struct {
	DataType string
}

func (sd SetDataType) String() string {
	return fmt.Sprintf(`SET DATA TYPE %s`, sd.DataType)
}

var _ Change = (*SetDataType)(nil)
