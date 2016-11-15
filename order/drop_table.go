package order

import (
	"fmt"
)

// Drop an existing table. This change occurs when you have a table name
// in the old schema with no match in the new schema.
type DropTable struct {
	SchemaName string
	TableName  string
}

func (x DropTable) String() string {
	return fmt.Sprintf(`DROP TABLE %s.%s`, x.SchemaName, x.TableName)
}

var _ Change = (*DropTable)(nil)
