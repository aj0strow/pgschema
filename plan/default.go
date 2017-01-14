package plan

import (
	"github.com/aj0strow/pgschema/db"
)

func setDefault(a, b *db.Column) bool {
	return !defaultEqual(a, b) && a.Default != ""
}

func dropDefault(a, b *db.Column) bool {
	return !defaultEqual(a, b) && b.Default != ""
}

func defaultEqual(a, b *db.Column) bool {
	return a.Default == b.Default || a.Default+"::"+a.DataType == b.Default
}
