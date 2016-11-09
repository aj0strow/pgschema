package ab

import (
	"github.com/aj0strow/pgschema/db"
)

type DatabaseMatch struct {
	SchemaMatches []SchemaMatch
}

func MatchDatabase(a, b db.DatabaseNode) DatabaseMatch {
	return DatabaseMatch{
		SchemaMatches: MatchSchemas(a.SchemaNodes, b.SchemaNodes),
	}
}
