package plan

import (
	"github.com/aj0strow/pgschema/db"
)

func setDataType(a, b *db.Column) bool {
	return a.DataType != b.DataType
}
