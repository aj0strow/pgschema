package pgschema

import (
	"github.com/aj0strow/pgschema/tree"
)

func planSchemaMatch(schema tree.SchemaMatch) []Change {
	var cs []Change
	tables := planTableMatches(schema.TableMatches)
	cs = append(cs, tables...)
	return cs
}

func planTableMatches(tableMatches []tree.TableMatch) []Change {
	var cs []Change
	for _, tableMatch := range tableMatches {
		tableChanges := planTableMatch(tableMatch)
		cs = append(cs, tableChanges...)
	}
	return cs
}

func planTableMatch(tableMatch tree.TableMatch) []Change {
	var cs []Change
	a, b := tableMatch.A, tableMatch.B
	if a == nil {
		cs = append(cs, DropTable{b.TableName})
	} else if b == nil {
		cs = append(cs, CreateTable{a.TableName})
	}
	if a != nil {
		columns := planColumnMatches(tableMatch.ColumnMatches)
		for _, change := range columns {
			cs = append(cs, AlterTable{a.TableName, change})
		}
	}
	return cs
}

func planColumnMatches(columnMatches []tree.ColumnMatch) []Change {
	var cs []Change
	for _, columnMatch := range columnMatches {
		changes := planColumnMatch(columnMatch)
		cs = append(cs, changes...)
	}
	return cs
}

func planColumnMatch(columnMatch tree.ColumnMatch) []Change {
	var cs []Change
	a, b := columnMatch.A, columnMatch.B
	if a == nil {
		cs = append(cs, DropColumn{b.ColumnName})
	} else if b == nil {
		cs = append(cs, AddColumn{a.ColumnName, a.DataType})
	} else if a.DataType != b.DataType {
		cs = append(cs, AlterColumn{a.ColumnName, SetDataType{a.DataType}})
	}
	return cs
}
