package order

import (
	"fmt"
)

// Alter an existing table. This change occurs when you have matching
// table names but the schema doesn't match.
type AlterTable struct {
	SchemaName string
	TableName  string
	Change     Change
}

func (x AlterTable) String() string {
	return fmt.Sprintf(`ALTER TABLE %s.%s %s`, x.SchemaName, x.TableName, x.Change)
}

var _ Change = (*AlterTable)(nil)
