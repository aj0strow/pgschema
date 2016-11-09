package plan

import (
	"github.com/aj0strow/pgschema/ab"
)

func DatabaseChanges(databaseMatch ab.DatabaseMatch) []Change {
	var cs []Change
	for _, schemaMatch := range databaseMatch.SchemaMatches {
		cs = append(cs, SchemaChanges(schemaMatch)...)
	}
	return cs
}

func SchemaChanges(schemaMatch ab.SchemaMatch) []Change {
	var cs []Change
	for _, tableMatch := range schemaMatch.TableMatches {
		cs = append(cs, TableChanges(tableMatch)...)
	}
	return cs
}
