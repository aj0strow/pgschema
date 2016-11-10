package plan

import (
	"fmt"
)

// Create a missing extension.
type CreateExtension struct {
	ExtName string
}

func (ce CreateExtension) String() string {
	return fmt.Sprintf(`CREATE EXTENSION "%s"`, ce.ExtName)
}

var _ Change = (*CreateExtension)(nil)

// Drop existing extension.
type DropExtension struct {
	ExtName string
}

func (de DropExtension) String() string {
	return fmt.Sprintf(`DROP EXTENSION "%s"`, de.ExtName)
}

var _ Change = (*DropExtension)(nil)
