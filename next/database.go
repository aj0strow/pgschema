package next

import (
	"github.com/aj0strow/pgschema/ab"
)

type UpdateDatabase struct {
	CreateExtensions []CreateExtension
	CreateSchemas    []CreateSchema
	UpdateSchemas    []UpdateSchema
}

func PlanUpdate(database ab.DatabaseMatch) UpdateDatabase {
	return updateDatabase(database)
}

func updateDatabase(database ab.DatabaseMatch) UpdateDatabase {
	return UpdateDatabase{
		CreateExtensions: createExtensions(database.ExtensionMatches),
		CreateSchemas:    createSchemas(database.SchemaMatches),
		UpdateSchemas:    updateSchemas(database.SchemaMatches),
	}
}
