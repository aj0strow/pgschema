package ab

import (
	"github.com/aj0strow/pgschema/info"
	"github.com/aj0strow/pgschema/tree"
)

// ColumnMatch is a combined column with new version A and old version B.
type ColumnMatch struct {
	A *info.Column
	B *info.Column
}

// MatchColumnNodes takes separate column node lists, and combines
// them by column name.
func MatchColumns(a, b []tree.ColumnNode) []ColumnMatch {
	var columnMatches []ColumnMatch
	fromA := map[string]bool{}
	for _, nodeA := range a {
		colA := nodeA.Column
		columnName := colA.ColumnName
		fromA[columnName] = true
		nodeB := findColumnNode(b, columnName)
		if nodeB != nil {
			colB := nodeB.Column
			columnMatches = append(columnMatches, ColumnMatch{
				A: &colA,
				B: &colB,
			})
		} else {
			columnMatches = append(columnMatches, ColumnMatch{
				A: &colA,
				B: nil,
			})
		}
	}
	for _, nodeB := range b {
		colB := nodeB.Column
		columnName := colB.ColumnName
		if !fromA[columnName] {
			columnMatches = append(columnMatches, ColumnMatch{
				A: nil,
				B: &colB,
			})
		}
	}
	return columnMatches
}

func findColumnNode(nodes []tree.ColumnNode, name string) *tree.ColumnNode {
	for _, node := range nodes {
		if node.Column.ColumnName == name {
			return &node
		}
	}
	return nil
}
