package next

import (
	"github.com/aj0strow/pgschema/db"
)

type SetDefault struct {
	Expression string
}

func setDefault(a, b *db.Column) *SetDefault {
	if a.Default != b.Default && a.Default != "" {
		return &SetDefault{
			Expression: a.Default,
		}
	}
	return nil
}

func dropDefault(a, b *db.Column) bool {
	return a.Default != b.Default && b.Default != ""
}
