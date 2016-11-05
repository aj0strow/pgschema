package pgschema

type Schema struct {
	SchemaName string
}

type Table struct {
	TableName string
	Columns   []Column
}

type Column struct {
	Table      *Table
	ColumnName string
	DataType   string
}
