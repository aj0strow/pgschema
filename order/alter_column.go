package order

import (
	"fmt"
)

// AlterColumn is a child change of AlterTable, and accepts
// a child column change.
type AlterColumn struct {
	ColumnName string
	Change     Change
}

func (x AlterColumn) String() string {
	return fmt.Sprintf(`ALTER COLUMN %s %s`, x.ColumnName, x.Change)
}

var _ Change = (*AlterColumn)(nil)
