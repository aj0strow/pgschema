package order

import (
	"fmt"
)

type SetDefault struct {
	Default string
}

func (x SetDefault) String() string {
	return fmt.Sprintf(`SET DEFAULT %s`, x.Default)
}

var _ Change = (*SetDefault)(nil)
