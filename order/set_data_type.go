package order

import (
	"bytes"
	"fmt"
)

// SetDataType is a child change of AlterColumn. It changes
// a data using a cast function.
type SetDataType struct {
	DataType         string
	NumericPrecision int
	NumericScale     int
	Using            string
}

func (x SetDataType) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "SET DATA TYPE %s", x.DataType)
	if x.DataType == "numeric" {
		fmt.Fprintf(&b, "(%d,%d)", x.NumericPrecision, x.NumericScale)
	}
	if x.Using != "" {
		fmt.Fprintf(&b, " USING %s", x.Using)
	}
	return b.String()
}

var _ Change = (*SetDataType)(nil)
