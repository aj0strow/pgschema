package plan

import (
	"github.com/aj0strow/pgschema/tree"
)

func DatabaseChanges(databaseMatch tree.DatabaseMatch) []Change {
	var cs []Change
	for _, schemaMatch := range databaseMatch.SchemaMatches {
		cs = append(cs, SchemaChanges(schemaMatch)...)
	}
	return cs
}

func SchemaChanges(schemaMatch tree.SchemaMatch) []Change {
	var cs []Change
	for _, tableMatch := range schemaMatch.TableMatches {
		cs = append(cs, TableChanges(tableMatch)...)
	}
	return cs
}
