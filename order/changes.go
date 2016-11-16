package order

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/plan"
)

func Changes(a, b db.DatabaseNode) []Change {
	return updateDatabase(plan.Update(ab.MatchDatabase(a, b)))
}

func updateDatabase(database plan.UpdateDatabase) []Change {
	var xs []Change
	for _, extension := range database.CreateExtensions {
		xs = append(xs, CreateExtension{
			ExtName: extension.ExtName,
		})
	}
	for _, schema := range database.UpdateSchemas {
		xs = append(xs, updateSchemaStruct(schema)...)
	}
	for _, schema := range database.CreateSchemas {
		xs = append(xs, createSchemaStruct(schema)...)
	}
	return xs
}

func updateSchemaStruct(schema plan.UpdateSchema) []Change {
	var xs []Change
	xs = append(xs, dropTableSlice(schema.Schema, schema.DropTables)...)
	xs = append(xs, alterTableSlice(schema.Schema, schema.AlterTables)...)
	xs = append(xs, createTableSlice(schema.Schema, schema.CreateTables)...)
	return xs
}

func createSchemaStruct(schema plan.CreateSchema) []Change {
	var xs []Change
	xs = append(xs, CreateSchema{
		SchemaName: schema.SchemaName,
	})
	xs = append(xs, createTableSlice(schema.Schema, schema.CreateTables)...)
	return xs
}

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
		xs = append(xs, dropIndexSlice(schema, table.Table, table.DropIndexes)...)

		for _, change := range alterTableColumns(table) {
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
