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
func MatchSchemas(a, b []db.SchemaNode) []SchemaMatch {
	var schemaMatches []SchemaMatch
	var fromA map[string]bool
	for _, nodeA := range a {
		schemaA := nodeA.Schema
		schemaName := schemaA.SchemaName
		fromA[schemaName] = true
		nodeB := findSchemaNode(b, schemaName)
		if nodeB != nil {
			schemaB := nodeB.Schema
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            &schemaA,
				B:            &schemaB,
				TableMatches: MatchTables(nodeA.TableNodes, nodeB.TableNodes),
			})
		} else {
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            &schemaA,
				B:            nil,
				TableMatches: MatchTables(nodeA.TableNodes, nil),
			})
		}
	}
	for _, nodeB := range b {
		schemaB := nodeB.Schema
		if !fromA[schemaB.SchemaName] {
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            nil,
				B:            &schemaB,
				TableMatches: MatchTables(nil, nodeB.TableNodes),
			})
		}
	}
	return schemaMatches
}

func findSchemaNode(nodes []db.SchemaNode, name string) *db.SchemaNode {
	for _, node := range nodes {
		if node.Schema.SchemaName == name {
			return &node
		}
	}
	return nil
}
