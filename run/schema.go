package run

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type CreateSchema struct {
	*db.Schema
	CreateTables []CreateTable
}

func createSchemas(schs []ab.SchemaMatch) []CreateSchema {
	var xs []CreateSchema
	for _, sch := range schs {
		if sch.B == nil {
			xs = append(xs, CreateSchema{
				Schema: sch.A,
			})
		}
	}
	return xs
}

type UpdateSchema struct {
	*db.Schema
	CreateTables []CreateTable
	DropTables   []DropTable
}

func updateSchemas(schemas []ab.SchemaMatch) []UpdateSchema {
	var xs []UpdateSchema
	for _, schema := range schemas {
		if schema.A != nil && schema.B != nil {
			xs = append(xs, UpdateSchema{
				Schema:       schema.A,
				CreateTables: createTables(schema.TableMatches),
				DropTables:   dropTables(schema.TableMatches),
			})
		}
	}
	return xs
}
