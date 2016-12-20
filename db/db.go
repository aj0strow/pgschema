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

type Index struct {
	TableName string
	IndexName string
	Exprs     []string
	Unique    bool
	Primary   bool
}
