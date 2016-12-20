package db

import (
	"testing"
)

func TestColumnErr(t *testing.T) {
	valid := []Column{
		{
			ColumnName: "email",
			DataType:   "citext",
		},
	}
	for _, c := range valid {
		if err := c.Err(); err != nil {
			t.Fatal(err)
		}
	}
	invalid := []Column{
		{
			ColumnName: "notype",
		},
		{
			DataType: "boolean",
		},
	}
	for _, c := range invalid {
		if err := c.Err(); err == nil {
			t.Errorf("invalid column: %# v\n", c)
		}
	}
}
