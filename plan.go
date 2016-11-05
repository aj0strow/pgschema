package pgschema

func PlanTables(a, b []Table) []Change {
	changes := []Change{}

	create := createTable(a, b)
	changes = append(changes, create...)

	drop := dropTable(a, b)
	changes = append(changes, drop...)

	return changes
}

func createTable(a, b []Table) []Change {
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

func dropTable(a, b []Table) []Change {
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
