package ab

import (
	"github.com/aj0strow/pgschema/db"
)

// ColumnMatch is a combined column with new version A and old version B.
type ColumnMatch struct {
	A *db.Column
	B *db.Column
}

// MatchColumnNodes takes separate column node lists, and combines
// them by column name.
func MatchColumns(a, b []*db.Column) []ColumnMatch {
	var columnMatches []ColumnMatch
	fromA := map[string]bool{}
	for _, colA := range a {
		columnName := colA.ColumnName
		fromA[columnName] = true
		colB := findColumn(b, columnName)
		if colB != nil {
			columnMatches = append(columnMatches, ColumnMatch{
				A: colA,
				B: colB,
			})
		} else {
			columnMatches = append(columnMatches, ColumnMatch{
				A: colA,
				B: nil,
			})
		}
	}
	for _, colB := range b {
		columnName := colB.ColumnName
		if !fromA[columnName] {
			columnMatches = append(columnMatches, ColumnMatch{
				A: nil,
				B: colB,
			})
		}
	}
	return columnMatches
}

func findColumn(cs []*db.Column, name string) *db.Column {
	for _, c := range cs {
		if c.ColumnName == name {
			return c
		}
	}
	return nil
}
