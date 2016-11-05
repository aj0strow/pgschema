package pgschema

func PlanTables(a, b []Table) []Change {
	changes := []Change{}

	create := createTables(a, b)
	changes = append(changes, create...)

	drop := dropTables(a, b)
	changes = append(changes, drop...)

	return changes
}

func createTables(a, b []Table) []Change {
	changes := []Change{}
	for _, t := range a {
		if !hasTable(b, t) {
			ct := CreateTable{
				TableName: t.TableName,
			}
			changes = append(changes, ct)
		}
	}
	return changes
}

func dropTables(a, b []Table) []Change {
	changes := []Change{}
	for _, t := range b {
		if !hasTable(a, t) {
			dt := DropTable{
				TableName: t.TableName,
			}
			changes = append(changes, dt)
		}
	}
	return changes
}

func hasTable(ts []Table, t Table) bool {
	for _, t2 := range ts {
		if t2.TableName == t.TableName {
			return true
		}
	}
	return false
}

func PlanColumns(a, b []Column) []Change {
	changes := []Change{}
	add := addColumns(a, b)
	changes = append(changes, add...)
	drop := dropColumns(a, b)
	changes = append(changes, drop...)
	return changes
}

func addColumns(a, b []Column) []Change {
	changes := []Change{}
	for _, c := range a {
		if !hasColumn(b, c) {
			changes = append(changes, AddColumn{
				TableName:  c.Table.TableName,
				ColumnName: c.ColumnName,
				DataType:   c.DataType,
			})
		}
	}
	return changes
}

func dropColumns(a, b []Column) []Change {
	changes := []Change{}
	for _, c := range b {
		if !hasColumn(a, c) {
			changes = append(changes, DropColumn{
				TableName:  c.Table.TableName,
				ColumnName: c.ColumnName,
			})
		}
	}
	return changes
}

func hasColumn(cs []Column, c Column) bool {
	for _, c2 := range cs {
		sameTable := c2.Table.TableName == c.Table.TableName
		sameName := c2.ColumnName == c.ColumnName
		if sameTable && sameName {
			return true
		}
	}
	return false
}
