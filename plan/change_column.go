package plan

import (
	"fmt"
)

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

// Change the column type and cast using a custom expression.
type CastDataType struct {
	SetDataType SetDataType
	Using       string
}

func (st CastDataType) String() string {
	return fmt.Sprintf("%s USING (%s)", st.SetDataType, st.Using)
}

var _ Change = (*CastDataType)(nil)

// Change the column data type.
type SetDataType struct {
	DataType string
}

func (sd SetDataType) String() string {
	return fmt.Sprintf(`SET DATA TYPE %s`, sd.DataType)
}

var _ Change = (*SetDataType)(nil)

// Set column to not null.
type SetNotNull struct{}

func (SetNotNull) String() string {
	return `SET NOT NULL`
}

var _ Change = (*SetNotNull)(nil)

// Drop not null constraint.
type DropNotNull struct{}

func (DropNotNull) String() string {
	return `DROP NOT NULL`
}

var _ Change = (*DropNotNull)(nil)

// Set column default expression value.
type SetDefault struct {
	Expression string
}

func (sd SetDefault) String() string {
	return fmt.Sprintf(`SET DEFAULT %s`, sd.Expression)
}

var _ Change = (*SetDefault)(nil)

// Drop column default.
type DropDefault struct{}

func (DropDefault) String() string {
	return `DROP DEFAULT`
}

var _ Change = (*DropDefault)(nil)
