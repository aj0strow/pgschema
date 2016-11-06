package pgschema

// Schema is a database catalog or schema.
type Schema struct {
	SchemaName string
}

// Table is a database table in the current search path.
type Table struct {
	TableName string
}

// Column is a database column belonging to a table.
type Column struct {
	ColumnName string
	DataType   string
}
