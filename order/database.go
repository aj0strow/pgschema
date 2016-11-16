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
