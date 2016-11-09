package ab

import (
	"github.com/aj0strow/pgschema/db"
)

type DatabaseMatch struct {
	ExtensionMatches []ExtensionMatch
	SchemaMatches    []SchemaMatch
}

func MatchDatabase(a, b db.DatabaseNode) DatabaseMatch {
	return DatabaseMatch{
		ExtensionMatches: MatchExtensions(a.ExtensionNodes, b.ExtensionNodes),
		SchemaMatches:    MatchSchemas(a.SchemaNodes, b.SchemaNodes),
	}
}
