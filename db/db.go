package db

type Database struct {
	Extensions []*Extension
	Schemas    []*Schema
}

type Extension struct {
	ExtName string
}

type Schema struct {
	SchemaName string
	Tables     []*Table
}

type Table struct {
	TableName string
	Indexes   []*Index
	Columns   []*Column
}

type Column struct {
	ColumnName       string
	DataType         string
	CastTypeUsing    string
	NotNull          bool
	Default          string
	NumericPrecision int
	NumericScale     int
	Array            bool
}

type Index struct {
	TableName string
	IndexName string
	Exprs     []string
	Unique    bool
	Primary   bool
}
