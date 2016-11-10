package plan

import (
	"bytes"
	"fmt"
)

type CreateIndex struct {
	IndexName string
	TableName string
	Exprs     []string
}

func (ci CreateIndex) String() string {
	var b bytes.Buffer
	b.WriteString("CREATE INDEX ")
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
