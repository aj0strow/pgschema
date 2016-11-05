package pgschema

import (
	"fmt"
)

// Change is a single atomic schema change.
type Change interface {
	String() string
}

// Create a new table. This change occurs when you have a table name in
// the new schema with no match in the old schema.
type CreateTable struct {
	TableName string
}

func (ct CreateTable) String() string {
	return fmt.Sprintf("CREATE TABLE %s", ct.TableName)
}

var _ Change = (*CreateTable)(nil)

// Drop an existing table. This change occurs when you have a table name
// in the old schema with no match in the new schema.
type DropTable struct {
	TableName string
}

func (dt DropTable) String() string {
	return fmt.Sprintf("DROP TABLE %s", dt.TableName)
}

var _ Change = (*DropTable)(nil)

// Add a new column to an existing table.
type AddColumn struct {
	TableName  string
	ColumnName string
	DataType   string
}

func (ac AddColumn) String() string {
	return fmt.Sprintf(`ALTER TABLE %s ADD COLUMN %s %s`, ac.TableName, ac.ColumnName, ac.DataType)
}

var _ Change = (*AddColumn)(nil)

// Drop existing column from existing table.
type DropColumn struct {
	TableName  string
	ColumnName string
}

func (dc DropColumn) String() string {
	return fmt.Sprintf(`ALTER TABLE %s DROP COLUMN %s RESTRICT`)
}

var _ Change = (*DropColumn)(nil)
