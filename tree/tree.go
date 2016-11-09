package tree

import (
	"github.com/aj0strow/pgschema/db"
)

type DatabaseNode struct {
	SchemaNodes []SchemaNode
}

type SchemaNode struct {
	Schema     db.Schema
	TableNodes []TableNode
}

type TableNode struct {
	Table       db.Table
	ColumnNodes []ColumnNode
}

type ColumnNode struct {
	Column db.Column
}
