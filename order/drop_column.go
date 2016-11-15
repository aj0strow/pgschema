package order

import (
	"fmt"
)

// DropColumn drops an existing column from existing table. It's
// a child change that requires a parent AlterTable.
type DropColumn struct {
	ColumnName string
}

func (x DropColumn) String() string {
	return fmt.Sprintf(`DROP COLUMN %s RESTRICT`, x.ColumnName)
}

var _ Change = (*DropColumn)(nil)
