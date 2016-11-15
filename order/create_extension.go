package order

import (
	"fmt"
)

// Create a missing extension.
type CreateExtension struct {
	ExtName string
}

func (x CreateExtension) String() string {
	return fmt.Sprintf(`CREATE EXTENSION "%s"`, x.ExtName)
}

var _ Change = (*CreateExtension)(nil)
