package temp

import (
	"fmt"
)

var counter = 0

func randSchemaName() string {
	counter += 1
	return fmt.Sprintf("db%d", counter)
}
