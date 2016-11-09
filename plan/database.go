package plan

import (
	"github.com/aj0strow/pgschema/ab"
)

func DatabaseChanges(databaseMatch ab.DatabaseMatch) []Change {
	var cs []Change
	for _, extMatch := range databaseMatch.ExtensionMatches {
		cs = append(cs, ExtensionChanges(extMatch)...)
	}
	for _, schemaMatch := range databaseMatch.SchemaMatches {
		cs = append(cs, SchemaChanges(schemaMatch)...)
	}
	return cs
}
