package db

type DatabaseNode struct {
	SchemaNodes    []SchemaNode
	ExtensionNodes []ExtensionNode
}

type ExtensionNode struct {
	Extension Extension
}

type SchemaNode struct {
	Schema     Schema
	TableNodes []TableNode
}

type TableNode struct {
	Table       Table
	ColumnNodes []ColumnNode
	IndexNodes  []IndexNode
}

type ColumnNode struct {
	Column Column
}

type IndexNode struct {
	Index Index
}
