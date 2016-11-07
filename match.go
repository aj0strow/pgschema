package pgschema

// SchemaMatch is a combined schema with the new version A and old version B.
type SchemaMatch struct {
	A            *Schema
	B            *Schema
	TableMatches []TableMatch
}

// TableMatch is a combined table with new version A and old version B.
type TableMatch struct {
	A             *Table
	B             *Table
	ColumnMatches []ColumnMatch
}

// ColumnMatch is a combined column with new version A and old version B.
type ColumnMatch struct {
	A *Column
	B *Column
}

/*
func MatchSchemaNodes(a, b []SchemaNode) []SchemaMatch {
	var schemaMatches []SchemaMatch
	var fromA map[string]bool
	for _, nodeA := range a {
		schemaA := &nodeA.Schema
		schemaName := schemaA.SchemaName
		fromA[schemaName] = true
		nodeB := findSchemaNode(b, schemaName)
		if nodeB != nil {
			schemaB := &nodeB.Schema
			tables := MatchTableNodes(nodeA.TableNodes, nodeB.TableNodes)
			schemaMatches = append(schemaMatches, SchemaMatch{
				SchemaName: schemaName,
				A: schemaA,
				B: schemaB,
				TableMatches: tables,
			})
		} else {
			tables = MatchTableNodes(nodeA.TableNodes, nil)
			schemaMatches = append(schemaMatches, SchemaMatch{
				TableName: tableName,
				A: schemaA,
				B: nil,
				TableMatches: tables,
			})
		}
	}
	for _, nodeB := range b {
		schemaB := &nodeB.Schema
		schemaName := schemaB.SchemaName
		if !fromA[schemaName] {
			schemaMatches = append(schemaMatches, SchemaMatch{
				SchemaName: schemaName,
				A: nil,
				B: schemaB,
				TableMatches: tables,
			})
		}
	}
	return schemaMatches
}

func findSchemaNode()
*/

// MatchTableNodes takes separate TableNode lists, and combines them
// by table name.
func MatchTableNodes(a, b []TableNode) []TableMatch {
	var tableMatches []TableMatch
	fromA := map[string]bool{}
	for _, nodeA := range a {
		tableA := &nodeA.Table
		tableName := tableA.TableName
		fromA[tableName] = true
		nodeB := findTableNode(b, tableName)
		if nodeB != nil {
			tableB := &nodeB.Table
			columns := MatchColumnNodes(nodeA.ColumnNodes, nodeB.ColumnNodes)
			tableMatches = append(tableMatches, TableMatch{
				A:             tableA,
				B:             tableB,
				ColumnMatches: columns,
			})
		} else {
			columns := MatchColumnNodes(nodeA.ColumnNodes, nil)
			tableMatches = append(tableMatches, TableMatch{
				A:             tableA,
				B:             nil,
				ColumnMatches: columns,
			})
		}
	}
	for _, nodeB := range b {
		tableB := &nodeB.Table
		tableName := tableB.TableName
		if !fromA[tableName] {
			columns := MatchColumnNodes(nil, nodeB.ColumnNodes)
			tableMatches = append(tableMatches, TableMatch{
				A:             nil,
				B:             tableB,
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
