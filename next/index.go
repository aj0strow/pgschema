package next

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type CreateIndex struct {
	Index *db.Index
}

func createIndexes(indexes []ab.IndexMatch) []CreateIndex {
	var xs []CreateIndex
	for _, index := range indexes {
		if index.B == nil {
			x := CreateIndex{
				Index: index.A,
			}
			xs = append(xs, x)
		}
	}
	return xs
}
