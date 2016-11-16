package order

import (
	"github.com/aj0strow/pgschema/plan"
)

func addColumnSlice(columns []plan.AddColumn) []Change {
	var xs []Change
	for _, column := range columns {
		xs = append(xs, AddColumn{
			ColumnName: column.ColumnName,
			DataType:   column.DataType,
			NotNull:    column.NotNull,
			Default:    column.Default,
		})
	}
	return xs
}

func alterColumnSlice(columns []plan.AlterColumn) []Change {
	var xs []Change
	for _, column := range columns {
		for _, change := range alterColumnStruct(column) {
			xs = append(xs, AlterColumn{
				ColumnName: column.ColumnName,
				Change:     change,
			})
		}
	}
	return xs
}

func alterColumnStruct(column plan.AlterColumn) []Change {
	var xs []Change
	if column.DropDefault {
		xs = append(xs, DropDefault)
	}
	if column.DropNotNull {
		xs = append(xs, DropNotNull)
	}
	if column.SetDataType {
		xs = append(xs, SetDataType{
			DataType: column.DataType,
			Using:    column.CastTypeUsing,
		})
	}
	if column.SetDefault {
		xs = append(xs, SetDefault{column.Default})
	}
	if column.SetNotNull {
		xs = append(xs, SetNotNull)
	}
	return xs
}

func dropColumnSlice(columns []plan.DropColumn) []Change {
	var xs []Change
	for _, column := range columns {
		xs = append(xs, DropColumn{
			ColumnName: column.ColumnName,
		})
	}
	return xs
}
