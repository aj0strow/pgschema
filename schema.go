package pgschema

type Schema struct {
	SchemaName string
	Tables     []Table
}

type Table struct {
	TableName string
	Columns   []Column
}

type Column struct {
	ColumnName string
	DataType   string
}
