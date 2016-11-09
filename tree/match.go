package tree

import (
	"github.com/aj0strow/pgschema/info"
)

type DatabaseMatch struct {
	SchemaMatches []SchemaMatch
}

// SchemaMatch is a combined schema with the new version A and old version B.
type SchemaMatch struct {
	A            *info.Schema
	B            *info.Schema
	TableMatches []TableMatch
}

// TableMatch is a combined table with new version A and old version B.
type TableMatch struct {
	A             *info.Table
	B             *info.Table
	ColumnMatches []ColumnMatch
}

// ColumnMatch is a combined column with new version A and old version B.
type ColumnMatch struct {
	A *info.Column
	B *info.Column
}

func Match(a, b DatabaseNode) DatabaseMatch {
	return DatabaseMatch{
		SchemaMatches: MatchSchemaNodes(a.SchemaNodes, b.SchemaNodes),
	}
}

// MatchSchemaNodes takes separate SchemaNode lists, and deep
// merges them into one combined SchemaMatch list.
func MatchSchemaNodes(a, b []SchemaNode) []SchemaMatch {
	var schemaMatches []SchemaMatch
	var fromA map[string]bool
	for _, nodeA := range a {
		schemaA := nodeA.Schema
		schemaName := schemaA.SchemaName
		fromA[schemaName] = true
		nodeB := findSchemaNode(b, schemaName)
		if nodeB != nil {
			schemaB := nodeB.Schema
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            &schemaA,
				B:            &schemaB,
				TableMatches: MatchTableNodes(nodeA.TableNodes, nodeB.TableNodes),
			})
		} else {
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            &schemaA,
				B:            nil,
				TableMatches: MatchTableNodes(nodeA.TableNodes, nil),
			})
		}
	}
	for _, nodeB := range b {
		schemaB := nodeB.Schema
		if !fromA[schemaB.SchemaName] {
			schemaMatches = append(schemaMatches, SchemaMatch{
				A:            nil,
				B:            &schemaB,
				TableMatches: MatchTableNodes(nil, nodeB.TableNodes),
			})
		}
	}
	return schemaMatches
}

func findSchemaNode(nodes []SchemaNode, name string) *SchemaNode {
	for _, node := range nodes {
		if node.Schema.SchemaName == name {
			return &node
		}
	}
	return nil
}

// MatchTableNodes takes separate TableNode lists, and deep merges them
// by table name into one combined TableMatch list.
func MatchTableNodes(a, b []TableNode) []TableMatch {
	var tableMatches []TableMatch
	fromA := map[string]bool{}
	for _, nodeA := range a {
		tableA := nodeA.Table
		tableName := tableA.TableName
		fromA[tableName] = true
		nodeB := findTableNode(b, tableName)
		if nodeB != nil {
			tableB := nodeB.Table
			columns := MatchColumnNodes(nodeA.ColumnNodes, nodeB.ColumnNodes)
			tableMatches = append(tableMatches, TableMatch{
				A:             &tableA,
				B:             &tableB,
				ColumnMatches: columns,
			})
		} else {
			columns := MatchColumnNodes(nodeA.ColumnNodes, nil)
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
			columns := MatchColumnNodes(nil, nodeB.ColumnNodes)
			tableMatches = append(tableMatches, TableMatch{
				A:             nil,
				B:             &tableB,
				ColumnMatches: columns,
			})
		}
	}
	return tableMatches
}

func findTableNode(nodes []TableNode, name string) *TableNode {
	for _, node := range nodes {
		if node.Table.TableName == name {
			return &node
		}
	}
	return nil
}

// MatchColumnNodes takes separate column node lists, and combines
// them by column name.
func MatchColumnNodes(a, b []ColumnNode) []ColumnMatch {
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

func findColumnNode(nodes []ColumnNode, name string) *ColumnNode {
	for _, node := range nodes {
		if node.Column.ColumnName == name {
			return &node
		}
	}
	return nil
}
