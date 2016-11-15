package order

import (
	"fmt"
)

type DropIndex struct {
	SchemaName string
	IndexName  string
}

func (x DropIndex) String() string {
	return fmt.Sprintf(`DROP INDEX %s.%s`, x.SchemaName, x.IndexName)
}

var _ Change = (*DropIndex)(nil)
