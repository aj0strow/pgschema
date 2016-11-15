package order

import (
	"github.com/aj0strow/pgschema/next"
)

type Change interface {
	String() string
}

func Changes(database next.UpdateDatabase) []Change {
	var xs []Change

	// Drop old tables.
	for _, updateSchema := range database.UpdateSchemas {
		for _, dropTable := range updateSchema.DropTables {
			x := DropTable{
				SchemaName: updateSchema.SchemaName,
				TableName:  dropTable.TableName,
			}
			xs = append(xs, x)
		}
	}

	// Alter existing tables.
	for _, updateSchema := range database.UpdateSchemas {
		for _, alterTable := range updateSchema.AlterTables {

			// Drop old indexes.
			for _, dropIndex := range alterTable.DropIndexes {
				if dropIndex.Primary {
					x := AlterTable{
						SchemaName: updateSchema.SchemaName,
						TableName:  alterTable.TableName,
						Change: DropConstraint{
							ConstraintName: dropIndex.IndexName,
						},
					}
					xs = append(xs, x)
				} else {
					x := DropIndex{
						SchemaName: updateSchema.SchemaName,
						IndexName:  dropIndex.IndexName,
					}
					xs = append(xs, x)
				}
			}

			// Drop old columns.
			for _, dropColumn := range alterTable.DropColumns {
				x := AlterTable{
					SchemaName: updateSchema.SchemaName,
					TableName:  alterTable.TableName,
					Change: DropColumn{
						ColumnName: dropColumn.ColumnName,
					},
				}
				xs = append(xs, x)
			}

			// Alter existing columns.
			for _, alterColumn := range alterTable.AlterColumns {
				for _, change := range alterColumnChanges(alterColumn) {
					x := AlterTable{
						SchemaName: updateSchema.SchemaName,
						TableName:  alterTable.TableName,
						Change:     change,
					}
					xs = append(xs, x)
				}
			}
		}
	}

	return xs
}

func alterColumnChanges(alterColumn next.AlterColumn) []Change {
	var xs []Change
	// Drop not null constraints.
	if alterColumn.DropNotNull {
		xs = append(xs, DropNotNull)
	}
	if alterColumn.SetNotNull {
		xs = append(xs, SetNotNull)
	}
	if alterColumn.SetDataType {
		xs = append(xs, SetDataType{
			DataType: alterColumn.DataType,
			Using:    alterColumn.CastTypeUsing,
		})
	}
	return xs
}
