package plan

import (
	"github.com/aj0strow/pgschema/ab"
)

func SchemaChanges(schemaMatch ab.SchemaMatch) []Change {
	var cs []Change
	a, b := schemaMatch.A, schemaMatch.B
	if a == nil {
		return cs
	}
	if b == nil {
		cs = append(cs, CreateSchema{SchemaName: a.SchemaName})
	}
	for _, tableMatch := range schemaMatch.TableMatches {
		cs = append(cs, TableChanges(tableMatch)...)
	}
	if len(cs) > 0 {
		cs = append([]Change{SetSearchPath{SchemaName: a.SchemaName}}, cs...)
	}
	return cs
}
