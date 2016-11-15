package order

import (
	"bytes"
)

type AddPrimaryKey struct {
	Columns []string
}

func (x AddPrimaryKey) String() string {
	var b bytes.Buffer
	b.WriteString("ADD PRIMARY KEY (")
	for i, name := range x.Columns {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(name)
	}
	b.WriteString(")")
	return b.String()
}

var _ Change = (*AddPrimaryKey)(nil)
