package next

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
)

type CreateTable struct {
	*db.Table
}

func createTables(tables []ab.TableMatch) []CreateTable {
	var xs []CreateTable
	for _, table := range tables {
		if table.B == nil {
			xs = append(xs, CreateTable{
				Table: table.A,
			})
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
}

func alterTables(tables []ab.TableMatch) []AlterTable {
	var xs []AlterTable
	for _, table := range tables {
		if table.A != nil && table.B != nil {
			xs = append(xs, AlterTable{
				Table: table.A,
			})
		}
	}
	return xs
}
