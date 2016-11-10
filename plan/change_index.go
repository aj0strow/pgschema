package plan

import (
	"bytes"
	"fmt"
)

type CreateIndex struct {
	TableName string
	IndexName string
	Exprs     []string
	Unique    bool
}

func (ci CreateIndex) String() string {
	var b bytes.Buffer
	b.WriteString("CREATE ")
	if ci.Unique {
		b.WriteString("UNIQUE ")
	}
	b.WriteString("INDEX ")
	b.WriteString(ci.IndexName)
	b.WriteString(" ON ")
	b.WriteString(ci.TableName)
	b.WriteString(" (")
	for i, expr := range ci.Exprs {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(expr)
	}
	b.WriteString(")")
	return b.String()
}

var _ Change = (*CreateIndex)(nil)

type DropIndex struct {
	IndexName string
}

func (di DropIndex) String() string {
	return fmt.Sprintf(`DROP INDEX %s`, di.IndexName)
}

var _ Change = (*DropIndex)(nil)

type AddPrimaryKey struct {
	ConstraintName string
	ColumnNames    []string
}

func (ap AddPrimaryKey) String() string {
	var b bytes.Buffer
	b.WriteString("ADD CONSTRAINT ")
	b.WriteString(ap.ConstraintName)
	b.WriteString(" PRIMARY KEY (")
	for i, name := range ap.ColumnNames {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(name)
	}
	b.WriteString(")")
	return b.String()
}

var _ Change = (*AddPrimaryKey)(nil)

type DropConstraint struct {
	ConstraintName string
}

func (dc DropConstraint) String() string {
	return fmt.Sprintf(`DROP CONSTRAINT %s`, dc.ConstraintName)
}

var _ Change = (*DropConstraint)(nil)
