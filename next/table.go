package next

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type CreateTable struct {
	*db.Table
	AddColumns    []AddColumn
	CreateIndexes []CreateIndex
}

func createTables(tables []ab.TableMatch) []CreateTable {
	var xs []CreateTable
	for _, table := range tables {
		if table.B == nil {
			x := CreateTable{
				Table:         table.A,
				AddColumns:    addColumns(table.ColumnMatches),
				CreateIndexes: createIndexes(table.IndexMatches),
			}
			xs = append(xs, x)
		}
	}
	return xs
}

type DropTable struct {
	*db.Table
}

func dropTables(tables []ab.TableMatch) []DropTable {
	var xs []DropTable
	for _, table := range tables {
		if table.A == nil {
			xs = append(xs, DropTable{
				Table: table.B,
			})
		}
	}
	return xs
}

type AlterTable struct {
	*db.Table
	CreateIndexes []CreateIndex
	DropIndexes   []DropIndex
}

func alterTables(tables []ab.TableMatch) []AlterTable {
	var xs []AlterTable
	for _, table := range tables {
		if table.A != nil && table.B != nil {
			x := AlterTable{
				Table:         table.A,
				CreateIndexes: createIndexes(table.IndexMatches),
				DropIndexes:   dropIndexes(table.IndexMatches),
			}
			xs = append(xs, x)
		}
	}
	return xs
}
