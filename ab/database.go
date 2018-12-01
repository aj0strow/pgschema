package ab

import (
	"github.com/aj0strow/pgschema/db"
)

type DatabaseMatch struct {
	ExtensionMatches []ExtensionMatch
	SchemaMatches    []SchemaMatch
}

func MatchDatabase(a, b *db.Database) DatabaseMatch {
	return DatabaseMatch{
		ExtensionMatches: MatchExtensions(a.Extensions, b.Extensions),
		SchemaMatches:    MatchSchemas(a.Schemas, b.Schemas),
	}
}
