package plan

import (
	"github.com/aj0strow/pgschema/ab"
)

func IndexChanges(indexMatch ab.IndexMatch) []Change {
	var cs []Change
	a, b := indexMatch.A, indexMatch.B
	if a == nil {
		return append(cs, DropIndex{b.IndexName})
	}
	return cs
}
