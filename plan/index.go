package plan

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type CreateIndex struct {
	*db.Index
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

type DropIndex struct {
	*db.Index
}

func dropIndexes(indexes []ab.IndexMatch) []DropIndex {
	var xs []DropIndex
	for _, index := range indexes {
		if index.A == nil {
			x := DropIndex{
				Index: index.B,
			}
			xs = append(xs, x)
		}
	}
	return xs
}
