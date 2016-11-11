package run

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type CreateSchema struct {
	Schema *db.Schema
}

type UpdateSchema struct {
	Schema *db.Schema
}

func createSchemas(schs []ab.SchemaMatch) []CreateSchema {
	var xs []CreateSchema
	for _, sch := range schs {
		if sch.B == nil {
			xs = append(xs, CreateSchema{sch.A})
		}
	}
	return xs
}
