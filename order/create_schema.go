package order

import (
	"fmt"
)

type CreateSchema struct {
	SchemaName string
}

func (x CreateSchema) String() string {
	return fmt.Sprintf(`CREATE SCHEMA %s`, x.SchemaName)
}

var _ Change = (*CreateSchema)(nil)
