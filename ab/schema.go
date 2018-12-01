package ab

import (
	"github.com/aj0strow/pgschema/db"
)

// SchemaMatch is a combined schema with the new version A and old version B.
type SchemaMatch struct {
	A            *db.Schema
	B            *db.Schema
	TableMatches []TableMatch
}

// MatchSchemaNodes takes separate SchemaNode lists, and deep
// merges them into one combined SchemaMatch list.
func MatchSchemas(a, b []*db.Schema) []SchemaMatch {
	var schemaMatches []SchemaMatch
	fromA := map[string]bool{}
	for _, schemaA := range a {
		schemaName := schemaA.SchemaName
		fromA[schemaName] = true
		schemaB := findSchema(b, schemaName)
		if schemaB != nil {
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            schemaA,
				B:            schemaB,
				TableMatches: MatchTables(schemaA.Tables, schemaB.Tables),
			})
		} else {
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            schemaA,
				B:            nil,
				TableMatches: MatchTables(schemaA.Tables, nil),
			})
		}
	}
	for _, schemaB := range b {
		if !fromA[schemaB.SchemaName] {
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            nil,
				B:            schemaB,
				TableMatches: MatchTables(nil, schemaB.Tables),
			})
		}
	}
	return schemaMatches
}

func findSchema(nodes []*db.Schema, name string) *db.Schema {
	for _, node := range nodes {
		if node.SchemaName == name {
			return node
		}
	}
	return nil
}
