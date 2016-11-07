package temp

import (
	"testing"
)

func TestRandSchemaName(t *testing.T) {
	a := randSchemaName()
	b := randSchemaName()
	if a == b {
		t.Fatal("schema names not random")
	}
}
