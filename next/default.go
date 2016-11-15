package next

import (
	"github.com/aj0strow/pgschema/db"
)

func setDefault(a, b *db.Column) bool {
	return a.Default != b.Default && a.Default != ""
}

func dropDefault(a, b *db.Column) bool {
	return a.Default != b.Default && b.Default != ""
}
