package next

import (
	"github.com/aj0strow/pgschema/db"
)

func setNotNull(a, b *db.Column) bool {
	return a.NotNull && !b.NotNull
}

func dropNotNull(a, b *db.Column) bool {
	return !a.NotNull && b.NotNull
}
