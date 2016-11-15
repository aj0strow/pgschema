package order

import (
	"bytes"
)

type WriteColumn struct {
	ColumnName string
	DataType   string
	NotNull    bool
	Default    string
}

func (x WriteColumn) String() string {
	var b bytes.Buffer
	b.WriteString(x.ColumnName)
	b.WriteString(" ")
	b.WriteString(x.DataType)
	if x.NotNull {
		b.WriteString(" NOT NULL")
	}
	if x.Default != "" {
		b.WriteString(" DEFAULT ")
		b.WriteString(x.Default)
	}
	return b.String()
}

var _ Change = (*WriteColumn)(nil)
