package order

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/plan"
)

func createTableSlice(schema *db.Schema, tables []plan.CreateTable) []Change {
	var xs []Change
	for _, table := range tables {
		xs = append(xs, CreateTable{
			SchemaName: schema.SchemaName,
			TableName:  table.TableName,
		})
		for _, change := range addColumnSlice(table.AddColumns) {
			xs = append(xs, AlterTable{
				SchemaName: schema.SchemaName,
				TableName:  table.TableName,
				Change:     change,
			})
		}
		xs = append(xs, createIndexSlice(schema, table.Table, table.CreateIndexes)...)
	}
	return xs
}

func alterTableSlice(schema *db.Schema, tables []plan.AlterTable) []Change {
	var xs []Change
	for _, table := range tables {
		xs = append(xs, alterTableStruct(schema, table)...)
	}
	return xs
}

func alterTableStruct(schema *db.Schema, table plan.AlterTable) []Change {
	var xs []Change
	xs = append(xs, dropIndexSlice(schema, table.Table, table.DropIndexes)...)
	for _, change := range alterTableColumns(table) {
		xs = append(xs, AlterTable{
			SchemaName: schema.SchemaName,
			TableName:  table.TableName,
			Change:     change,
		})
	}
	xs = append(xs, createIndexSlice(schema, table.Table, table.CreateIndexes)...)
	return xs
}

func alterTableColumns(table plan.AlterTable) []Change {
	var xs []Change
	xs = append(xs, dropColumnSlice(table.DropColumns)...)
	xs = append(xs, alterColumnSlice(table.AlterColumns)...)
	xs = append(xs, addColumnSlice(table.AddColumns)...)
	return xs
}

func dropTableSlice(schema *db.Schema, tables []plan.DropTable) []Change {
	var xs []Change
	for _, table := range tables {
		xs = append(xs, DropTable{
			SchemaName: schema.SchemaName,
			TableName:  table.TableName,
		})
	}
	return xs
}
