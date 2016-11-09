package ab

import (
	"github.com/aj0strow/pgschema/tree"
)

type DatabaseMatch struct {
	SchemaMatches []SchemaMatch
}

func MatchDatabase(a, b tree.DatabaseNode) DatabaseMatch {
	return DatabaseMatch{
		SchemaMatches: MatchSchemas(a.SchemaNodes, b.SchemaNodes),
	}
}
