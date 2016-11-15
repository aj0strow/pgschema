package order

import (
	"fmt"
)

// SetDataType is a child change of AlterColumn. It changes
// a data using a cast function.
type SetDataType struct {
	DataType string
	Using    string
}

func (x SetDataType) String() string {
	setDataType := fmt.Sprintf(`SET DATA TYPE %s`, x.DataType)
	if x.Using == "" {
		return setDataType
	}
	return fmt.Sprintf(`%s USING %s`, setDataType, x.Using)
}

var _ Change = (*SetDataType)(nil)
