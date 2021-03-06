package plan

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type CreateSchema struct {
	*db.Schema
	CreateTables []CreateTable
}

func createSchemas(schemas []ab.SchemaMatch) []CreateSchema {
	var xs []CreateSchema
	for _, schema := range schemas {
		if schema.B == nil {
			x := CreateSchema{
				Schema:       schema.A,
				CreateTables: createTables(schema.TableMatches),
			}
			xs = append(xs, x)
		}
	}
	return xs
}

type UpdateSchema struct {
	*db.Schema
	CreateTables []CreateTable
	AlterTables  []AlterTable
	DropTables   []DropTable
}

func updateSchemas(schemas []ab.SchemaMatch) []UpdateSchema {
	var xs []UpdateSchema
	for _, schema := range schemas {
		if schema.A != nil && schema.B != nil {
			x := UpdateSchema{
				Schema:       schema.A,
				CreateTables: createTables(schema.TableMatches),
				AlterTables:  alterTables(schema.TableMatches),
				DropTables:   dropTables(schema.TableMatches),
			}
			xs = append(xs, x)
		}
	}
	return xs
}
