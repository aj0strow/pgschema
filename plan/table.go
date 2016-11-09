package plan

import (
	"github.com/aj0strow/pgschema/ab"
)

func TableChanges(tableMatch ab.TableMatch) []Change {
	var cs []Change
	a, b := tableMatch.A, tableMatch.B
	if a == nil {
		return append(cs, DropTable{b.TableName})
	}
	if b == nil {
		cs = append(cs, CreateTable{a.TableName})
	}
	for _, columnMatch := range tableMatch.ColumnMatches {
		for _, change := range ColumnChanges(columnMatch) {
			cs = append(cs, AlterTable{a.TableName, change})
		}
	}
	return cs
}
