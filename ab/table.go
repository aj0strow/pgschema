package ab

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/tree"
)

// TableMatch is a combined table with new version A and old version B.
type TableMatch struct {
	A             *db.Table
	B             *db.Table
	ColumnMatches []ColumnMatch
}

// MatchTableNodes takes separate TableNode lists, and deep merges them
// by table name into one combined TableMatch list.
func MatchTables(a, b []tree.TableNode) []TableMatch {
	var tableMatches []TableMatch
	fromA := map[string]bool{}
	for _, nodeA := range a {
		tableA := nodeA.Table
		tableName := tableA.TableName
		fromA[tableName] = true
		nodeB := findTableNode(b, tableName)
		if nodeB != nil {
			tableB := nodeB.Table
			columns := MatchColumns(nodeA.ColumnNodes, nodeB.ColumnNodes)
			tableMatches = append(tableMatches, TableMatch{
				A:             &tableA,
				B:             &tableB,
				ColumnMatches: columns,
			})
		} else {
			columns := MatchColumns(nodeA.ColumnNodes, nil)
			tableMatches = append(tableMatches, TableMatch{
				A:             &tableA,
				B:             nil,
				ColumnMatches: columns,
			})
		}
	}
	for _, nodeB := range b {
		tableB := nodeB.Table
		tableName := tableB.TableName
		if !fromA[tableName] {
			columns := MatchColumns(nil, nodeB.ColumnNodes)
			tableMatches = append(tableMatches, TableMatch{
				A:             nil,
				B:             &tableB,
				ColumnMatches: columns,
			})
		}
	}
	return tableMatches
}

func findTableNode(nodes []tree.TableNode, name string) *tree.TableNode {
	for _, node := range nodes {
		if node.Table.TableName == name {
			return &node
		}
	}
	return nil
}
