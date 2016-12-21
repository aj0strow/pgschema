package db

import (
	"fmt"
)

type Column struct {
	ColumnName    string
	DataType      string
	CastTypeUsing string
	NotNull       bool
	Default       string
}

func (c *Column) Err() error {
	if c.ColumnName == "" {
		return fmt.Errorf("pgschema: column name can't be empty")
	}
	if c.DataType == "" {
		return fmt.Errorf("pgschema: column data type can't be empty")
	}
	return nil
}

type ColumnNode struct {
	Column Column
}

func (cn *ColumnNode) Err() error {
	return cn.Column.Err()
}
