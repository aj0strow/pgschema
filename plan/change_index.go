package plan

import (
	"fmt"
)

type DropIndex struct {
	IndexName string
}

func (di DropIndex) String() string {
	return fmt.Sprintf(`DROP INDEX %s`, di.IndexName)
}

var _ Change = (*DropIndex)(nil)
