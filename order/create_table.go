package order

import (
	"fmt"
)

type CreateTable struct {
	SchemaName string
	TableName  string
}

func (x CreateTable) String() string {
	return fmt.Sprintf(`CREATE TABLE %s.%s ()`, x.SchemaName, x.TableName)
}

var _ Change = (*CreateTable)(nil)
