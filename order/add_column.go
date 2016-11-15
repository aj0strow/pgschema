package order

import (
	"bytes"
)

type AddColumn struct {
	ColumnName string
	DataType   string
	NotNull    bool
	Default    string
}

func (x AddColumn) String() string {
	var b bytes.Buffer
	b.WriteString("ADD COLUMN ")
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

var _ Change = (*AddColumn)(nil)
