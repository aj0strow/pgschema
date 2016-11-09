package db

type DatabaseNode struct {
	SchemaNodes []SchemaNode
}

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
