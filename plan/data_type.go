package plan

import (
	"github.com/aj0strow/pgschema/db"
)

func setDataType(a, b *db.Column) bool {
	return getCanonicalType(a.DataType) != getCanonicalType(b.DataType)
}

var typeMap = map[string]string{
	"timestamp":   "timestamp without time zone",
	"timestamptz": "timestamp with time zone",
}

func getCanonicalType(dt string) string {
	if ct, ok := typeMap[dt]; ok {
		return ct
	}
	return dt
}
