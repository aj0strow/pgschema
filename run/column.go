package run

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
			xs = append(xs, AddColumn{
				Column: column.A,
			})
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
			xs = append(xs, DropColumn{
				Column: column.B,
			})
		}
	}
	return xs
}
