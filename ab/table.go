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
func MatchTables(a, b []*db.Table) []TableMatch {
	var tableMatches []TableMatch
	fromA := map[string]bool{}
	for _, tableA := range a {
		tableName := tableA.TableName
		fromA[tableName] = true
		tableB := findTable(b, tableName)
		if tableB != nil {
			tableMatches = append(tableMatches, TableMatch{
				A:             tableA,
				B:             tableB,
				ColumnMatches: MatchColumns(tableA.Columns, tableB.Columns),
				IndexMatches:  MatchIndexes(tableA.Indexes, tableB.Indexes),
			})
		} else {
			tableMatches = append(tableMatches, TableMatch{
				A:             tableA,
				B:             nil,
				ColumnMatches: MatchColumns(tableA.Columns, nil),
				IndexMatches:  MatchIndexes(tableA.Indexes, nil),
			})
		}
	}
	for _, tableB := range b {
		if !fromA[tableB.TableName] {
			tableMatches = append(tableMatches, TableMatch{
				A:             nil,
				B:             tableB,
				ColumnMatches: MatchColumns(nil, tableB.Columns),
				IndexMatches:  MatchIndexes(nil, tableB.Indexes),
			})
		}
	}
	return tableMatches
}

func findTable(nodes []*db.Table, name string) *db.Table {
	for _, node := range nodes {
		if node.TableName == name {
			return node
		}
	}
	return nil
}
