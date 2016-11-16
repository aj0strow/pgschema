package order

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/plan"
)

func createIndexSlice(schema *db.Schema, table *db.Table, indexes []plan.CreateIndex) []Change {
	var xs []Change
	for _, index := range indexes {
		xs = append(xs, createIndexStruct(schema, table, index.Index))
	}
	return xs
}

func createIndexStruct(schema *db.Schema, table *db.Table, index *db.Index) Change {
	if index.Primary {
		return AlterTable{
			SchemaName: schema.SchemaName,
			TableName:  table.TableName,
			Change: AddPrimaryKey{
				Columns: index.Exprs,
			},
		}
	} else {
		return CreateIndex{
			SchemaName: schema.SchemaName,
			TableName:  table.TableName,
			IndexName:  index.IndexName,
			Exprs:      index.Exprs,
			Unique:     index.Unique,
		}
	}
}

func dropIndexSlice(schema *db.Schema, table *db.Table, indexes []plan.DropIndex) []Change {
	var xs []Change
	for _, index := range indexes {
		xs = append(xs, dropIndexStruct(schema, table, index.Index))
	}
	return xs
}

func dropIndexStruct(schema *db.Schema, table *db.Table, index *db.Index) Change {
	if index.Primary {
		return AlterTable{
			SchemaName: schema.SchemaName,
			TableName:  table.TableName,
			Change: DropConstraint{
				ConstraintName: index.IndexName,
			},
		}
	} else {
		return DropIndex{
			SchemaName: schema.SchemaName,
			IndexName:  index.IndexName,
		}
	}
}
