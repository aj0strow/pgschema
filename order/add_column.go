package order

import (
	"fmt"
)

type AddColumn WriteColumn

func (x AddColumn) String() string {
	return fmt.Sprintf(`ADD COLUMN %s`, WriteColumn(x))
}

var _ Change = (*AddColumn)(nil)
