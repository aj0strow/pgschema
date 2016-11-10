package plan

import (
	"strings"

	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

func ColumnChanges(columnMatch ab.ColumnMatch) []Change {
	var cs []Change
	a, b := columnMatch.A, columnMatch.B
	if a == nil {
		return append(cs, DropColumn{b.ColumnName})
	}
	if b == nil {
		cs = append(cs, AddColumn{a.ColumnName, a.DataType})
		for _, change := range createColumn(a) {
			cs = append(cs, AlterColumn{a.ColumnName, change})
		}
		return cs
	}
	for _, change := range alterColumn(a, b) {
		cs = append(cs, AlterColumn{a.ColumnName, change})
	}
	return cs
}

func createColumn(a *db.Column) []Change {
	var cs []Change
	if a.NotNull {
		cs = append(cs, SetNotNull{})
	}
	return cs
}

func alterColumn(a, b *db.Column) []Change {
	var cs []Change
	if a.DataType != b.DataType {
		if a.CastTypeUsing == "" {
			cs = append(cs, SetDataType{a.DataType})
		} else {
			cs = append(cs, CastDataType{
				SetDataType: SetDataType{a.DataType},
				Using:       strings.Replace(a.CastTypeUsing, "$name", a.ColumnName, -1),
			})
		}
	}
	if a.NotNull != b.NotNull {
		if a.NotNull {
			cs = append(cs, SetNotNull{})
		} else {
			cs = append(cs, DropNotNull{})
		}
	}
	return cs
}
