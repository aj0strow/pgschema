package order

import (
	"github.com/aj0strow/pgschema/next"
)

type Change interface {
	String() string
}

func Changes(database next.UpdateDatabase) []Change {
	var xs []Change

	for _, updateSchema := range database.UpdateSchemas {
		// Drop old tables.
		for _, dropTable := range updateSchema.DropTables {
			x := DropTable{
				SchemaName: updateSchema.SchemaName,
				TableName:  dropTable.TableName,
			}
			xs = append(xs, x)
		}

		// Alter existing tables.
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

			// Add new columns.
			for _, addColumn := range alterTable.AddColumns {
				x := AlterTable{
					SchemaName: updateSchema.SchemaName,
					TableName:  alterTable.TableName,
					Change: AddColumn{
						ColumnName: addColumn.ColumnName,
						DataType:   addColumn.DataType,
						NotNull:    addColumn.NotNull,
						Default:    addColumn.Default,
					},
				}
				xs = append(xs, x)
			}

			// Create new indexes.
			for _, createIndex := range alterTable.CreateIndexes {
				if createIndex.Primary {
					x := AlterTable{
						SchemaName: updateSchema.SchemaName,
						TableName:  alterTable.TableName,
						Change: AddPrimaryKey{
							Columns: createIndex.Exprs,
						},
					}
					xs = append(xs, x)
				} else {
					x := CreateIndex{
						SchemaName: updateSchema.SchemaName,
						IndexName:  createIndex.IndexName,
						Exprs:      createIndex.Exprs,
						Unique:     createIndex.Unique,
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
