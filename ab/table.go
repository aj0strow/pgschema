package ab

import (
	"github.com/aj0strow/pgschema/db"
)

// TableMatch is a combined table with new version A and old version B.
type TableMatch struct {
	A             *db.Table
	B             *db.Table
	ColumnMatches []ColumnMatch
	IndexMatches  []IndexMatch
}

// MatchTableNodes takes separate TableNode lists, and deep merges them
// by table name into one combined TableMatch list.
func MatchTables(a, b []db.TableNode) []TableMatch {
	var tableMatches []TableMatch
	fromA := map[string]bool{}
	for _, nodeA := range a {
		tableA := nodeA.Table
		tableName := tableA.TableName
		fromA[tableName] = true
		nodeB := findTableNode(b, tableName)
		if nodeB != nil {
			tableB := nodeB.Table
			tableMatches = append(tableMatches, TableMatch{
				A:             &tableA,
				B:             &tableB,
				ColumnMatches: MatchColumns(nodeA.ColumnNodes, nodeB.ColumnNodes),
				IndexMatches:  MatchIndexes(nodeA.IndexNodes, nodeB.IndexNodes),
			})
		} else {
			tableMatches = append(tableMatches, TableMatch{
				A:             &tableA,
				B:             nil,
				ColumnMatches: MatchColumns(nodeA.ColumnNodes, nil),
				IndexMatches:  MatchIndexes(nodeA.IndexNodes, nil),
			})
		}
	}
	for _, nodeB := range b {
		tableB := nodeB.Table
		tableName := tableB.TableName
		if !fromA[tableName] {
			tableMatches = append(tableMatches, TableMatch{
				A:             nil,
				B:             &tableB,
				ColumnMatches: MatchColumns(nil, nodeB.ColumnNodes),
				IndexMatches:  MatchIndexes(nil, nodeB.IndexNodes),
			})
		}
	}
	return tableMatches
}

func findTableNode(nodes []db.TableNode, name string) *db.TableNode {
	for _, node := range nodes {
		if node.Table.TableName == name {
			return &node
		}
	}
	return nil
}
