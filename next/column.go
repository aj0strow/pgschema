package next

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type AddColumn struct {
	Column *db.Column
}

func addColumns(columns []ab.ColumnMatch) []AddColumn {
	var xs []AddColumn
	for _, column := range columns {
		if column.B == nil {
			x := AddColumn{
				Column: column.A,
			}
			xs = append(xs, x)
		}
	}
	return xs
}

type DropColumn struct {
	Column *db.Column
}

func dropColumns(columns []ab.ColumnMatch) []DropColumn {
	var xs []DropColumn
	for _, column := range columns {
		if column.A == nil {
			x := DropColumn{
				Column: column.B,
			}
			xs = append(xs, x)
		}
	}
	return xs
}

type AlterColumn struct {
	Column      *db.Column
	SetNotNull  bool
	DropNotNull bool
}

func alterColumns(columns []ab.ColumnMatch) []AlterColumn {
	var xs []AlterColumn
	for _, column := range columns {
		if column.A != nil && column.B != nil {
			x := alterColumn(column.A, column.B)
			xs = append(xs, x)
		}
	}
	return xs
}

func alterColumn(a, b *db.Column) AlterColumn {
	return AlterColumn{
		Column:      a,
		SetNotNull:  setNotNull(a, b),
		DropNotNull: dropNotNull(a, b),
	}
}
