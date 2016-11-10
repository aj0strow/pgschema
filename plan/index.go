package plan

import (
	"reflect"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

func IndexChanges(indexMatch ab.IndexMatch) []Change {
	var cs []Change
	a, b := indexMatch.A, indexMatch.B
	if a == nil {
		return append(cs, dropIndex(b))
	}
	if b == nil {
		return append(cs, createIndex(a))
	}
	if !ixEqual(a, b) {
		return append(cs, dropIndex(b), createIndex(a))
	}
	return cs
}

func ixEqual(a, b *db.Index) bool {
	ok := a.TableName == b.TableName && a.Unique == b.Unique && a.Primary == b.Primary
	if a.Primary && b.Primary {
		return ok && reflect.DeepEqual(a.Exprs, b.Exprs)
	} else {
		return ok
	}
}

func createIndex(ix *db.Index) Change {
	if ix.Primary {
		return AlterTable{ix.TableName, AddPrimaryKey{ix.IndexName, ix.Exprs}}
	} else {
		return CreateIndex{
			TableName: ix.TableName,
			IndexName: ix.IndexName,
			Exprs:     ix.Exprs,
			Unique:    ix.Unique,
		}
	}
}

func dropIndex(ix *db.Index) Change {
	if ix.Primary {
		return AlterTable{ix.TableName, DropConstraint{ix.IndexName}}
	} else {
		return DropIndex{ix.IndexName}
	}
}
