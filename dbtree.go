package pgschema

type SchemaNode struct {
	Schema     Schema
	TableNodes []TableNode
}

type TableNode struct {
	Table       Table
	ColumnNodes []ColumnNode
}

type ColumnNode struct {
	Column Column
}

func LoadSchemaNode(pg PG, schema Schema) (SchemaNode, error) {
	tables, err := LoadTables(pg, schema.SchemaName)
	if err != nil {
		return SchemaNode{}, err
	}
	var tableNodes []TableNode
	for _, table := range tables {
		tableNode, err := LoadTableNode(pg, schema, table)
		if err != nil {
			return SchemaNode{}, err
		}
		tableNodes = append(tableNodes, tableNode)
	}
	schemaNode := SchemaNode{
		Schema:     schema,
		TableNodes: tableNodes,
	}
	return schemaNode, nil
}

func LoadTableNode(pg PG, schema Schema, table Table) (TableNode, error) {
	columns, err := LoadColumns(pg, schema.SchemaName, table.TableName)
	if err != nil {
		return TableNode{}, err
	}
	var columnNodes []ColumnNode
	for _, column := range columns {
		columnNodes = append(columnNodes, ColumnNode{
			Column: column,
		})
	}
	tableNode := TableNode{
		Table:       table,
		ColumnNodes: columnNodes,
	}
	return tableNode, nil
}
