package plan

import (
	"github.com/aj0strow/pgschema/ab"
)

func ColumnChanges(columnMatch ab.ColumnMatch) []Change {
	var cs []Change
	a, b := columnMatch.A, columnMatch.B
	if a == nil {
		return append(cs, DropColumn{b.ColumnName})
	}
	if b == nil {
		cs = append(cs, AddColumn{a.ColumnName, a.DataType})
	} else if a.DataType != b.DataType {
		cs = append(cs, AlterColumn{a.ColumnName, SetDataType{a.DataType}})
	}
	return cs
}
