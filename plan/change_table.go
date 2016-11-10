package plan

import (
	"fmt"
)

// Create a new table. This change occurs when you have a table name in
// the new schema with no match in the old schema.
type CreateTable struct {
	TableName string
}

func (ct CreateTable) String() string {
	return fmt.Sprintf(`CREATE TABLE %s ()`, ct.TableName)
}

var _ Change = (*CreateTable)(nil)

// Drop an existing table. This change occurs when you have a table name
// in the old schema with no match in the new schema.
type DropTable struct {
	TableName string
}

func (dt DropTable) String() string {
	return fmt.Sprintf(`DROP TABLE %s`, dt.TableName)
}

var _ Change = (*DropTable)(nil)

// Alter an existing table. This change occurs when you have matching
// table names but the schema doesn't match.
type AlterTable struct {
	TableName string
	Change    Change
}

func (at AlterTable) String() string {
	return fmt.Sprintf(`ALTER TABLE %s %s`, at.TableName, at.Change)
}

var _ Change = (*AlterTable)(nil)
