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
	if b == nil {
		return append(cs, CreateIndex{
			IndexName: a.IndexName,
			TableName: a.TableName,
			Exprs:     a.Exprs,
		})
	}
	return cs
}
