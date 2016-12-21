package plan

import (
	"github.com/aj0strow/pgschema/db"
)

func setDataType(a, b *db.Column) bool {
	return !isSameType(a, b)
}

func isSameType(a, b *db.Column) bool {
	if a.DataType == "numeric" && b.DataType == "numeric" {
		return a.NumericPrecision == b.NumericPrecision && a.NumericScale == b.NumericScale
	}
	return getCanonicalType(a.DataType) == getCanonicalType(b.DataType)
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
