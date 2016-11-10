package db

type Extension struct {
	ExtName string
}

type Schema struct {
	SchemaName string
}

type Table struct {
	TableName string
}

type Column struct {
	ColumnName    string
	DataType      string
	CastTypeUsing string
	NotNull       bool
}

type Index struct {
	IndexName string
	Exprs     []string
}
