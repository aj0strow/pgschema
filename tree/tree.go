package tree

import (
	"github.com/aj0strow/pgschema/db"
	"github.com/aj0strow/pgschema/info"
)

type DatabaseNode struct {
	SchemaNodes []SchemaNode
}

type SchemaNode struct {
	Schema     db.Schema
	TableNodes []TableNode
}

type TableNode struct {
	Table       info.Table
	ColumnNodes []ColumnNode
}

type ColumnNode struct {
	Column info.Column
}
