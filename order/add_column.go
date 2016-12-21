package order

import (
	"bytes"
	"fmt"
)

type AddColumn struct {
	ColumnName       string
	DataType         string
	NumericPrecision int
	NumericScale     int
	NotNull          bool
	Default          string
}

func (x AddColumn) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ADD COLUMN %s %s", x.ColumnName, x.DataType)
	if x.DataType == "numeric" {
		fmt.Fprintf(&b, "(%d,%d)", x.NumericPrecision, x.NumericScale)
	}
	if x.NotNull {
		fmt.Fprintf(&b, " NOT NULL")
	}
	if x.Default != "" {
		fmt.Fprintf(&b, " DEFAULT %s", x.Default)
	}
	return b.String()
}

var _ Change = (*AddColumn)(nil)
