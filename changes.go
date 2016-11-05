package pgschema

import (
	"fmt"
)

type Change interface {
	String() string
}

type CreateTable struct {
	TableName string
}

func (ct CreateTable) String() string {
	return fmt.Sprintf("CREATE TABLE %s", ct.TableName)
}

var _ Change = (*CreateTable)(nil)

type DropTable struct {
	TableName string
}

func (dt DropTable) String() string {
	return fmt.Sprintf("DROP TABLE %s", dt.TableName)
}

var _ Change = (*DropTable)(nil)
