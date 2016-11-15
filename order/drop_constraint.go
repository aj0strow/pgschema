package order

import (
	"fmt"
)

// DropConstraint is a nested change that requires a parent
// AlterTable change.
type DropConstraint struct {
	ConstraintName string
}

func (x DropConstraint) String() string {
	return fmt.Sprintf(`DROP CONSTRAINT %s`, x.ConstraintName)
}

var _ Change = (*DropConstraint)(nil)
