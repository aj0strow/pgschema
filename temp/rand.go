package temp

import (
	"math/rand"
)

const chars = "abcdefghijklmnopqrstwxyz"

func init() {
	rand.Seed(100)
}

// randSchemaName generates a 10-character lowercase schema name.
func randSchemaName() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
