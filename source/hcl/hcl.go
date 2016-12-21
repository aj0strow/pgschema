// The hcl package parses an input configuration file to create a database db.
package hcl

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/aj0strow/pgschema/db"
	"github.com/hashicorp/hcl"
)

type Database struct {
	Extension map[string]Extension `hcl:"extension"`
	Schema    map[string]Schema    `hcl:"schema"`
}

type Extension struct{}

type Schema struct {
	Table map[string]Table `hcl:"table"`
}

type Table struct {
	PrimaryKey []string          `hcl:"primary_key"`
	Column     map[string]Column `hcl:"column"`
	Index      map[string]Index  `hcl:"index"`
}

type Column struct {
	Type          string `hcl:"type"`
	NotNull       bool   `hcl:"not_null"`
	CastTypeUsing string `hcl:"cast_type_using"`
	Default       string `hcl:"default"`
	PrimaryKey    bool   `hcl:"primary_key"`
}

type Index struct {
	On     []string `hcl:"on"`
	Unique bool     `hcl:"unique"`
}

func ParseBytes(bs []byte) (db.DatabaseNode, error) {
	var database Database
	err := hcl.Unmarshal(bs, &database)
	if err != nil {
		return db.DatabaseNode{}, err
	}
	databaseNode := convertDatabase(database)
	return databaseNode, nil
}

func convertDatabase(v Database) db.DatabaseNode {
	var exts []db.ExtensionNode
	for ek := range v.Extension {
		exts = append(exts, convertExtension(ek))
	}
	var schemas []db.SchemaNode
	for sk, sv := range v.Schema {
		schemas = append(schemas, convertSchema(sk, sv))
	}
	return db.DatabaseNode{
		ExtensionNodes: exts,
		SchemaNodes:    schemas,
	}
}

func convertExtension(k string) db.ExtensionNode {
	return db.ExtensionNode{
		Extension: db.Extension{
			ExtName: k,
		},
	}
}

func convertSchema(k string, v Schema) db.SchemaNode {
	var tables []db.TableNode
	for tableName, tv := range v.Table {
		tables = append(tables, convertTable(tableName, tv))
	}
	return db.SchemaNode{
		Schema: db.Schema{
			SchemaName: k,
		},
		TableNodes: tables,
	}
}

func convertTable(tableName string, v Table) db.TableNode {
	var (
		columns []db.ColumnNode
		indexes []db.IndexNode
	)
	for columnName, c := range v.Column {
		columns = append(columns, convertColumn(columnName, c))
		if c.PrimaryKey {
			indexes = append(indexes, newPrimaryKey(tableName, []string{columnName}))
		}
	}
	if len(v.PrimaryKey) > 0 {
		indexes = append(indexes, newPrimaryKey(tableName, v.PrimaryKey))
	}
	for indexName, ix := range v.Index {
		indexes = append(indexes, convertIndex(tableName, indexName, ix))
	}
	return db.TableNode{
		Table: db.Table{
			TableName: tableName,
		},
		ColumnNodes: columns,
		IndexNodes:  indexes,
	}
}

var numericRe = regexp.MustCompile(`numeric\((\d+),(\d+)\)`)

func convertColumn(k string, v Column) db.ColumnNode {
	c := db.Column{
		ColumnName:    k,
		DataType:      v.Type,
		CastTypeUsing: v.CastTypeUsing,
		NotNull:       v.NotNull || v.PrimaryKey,
		Default:       v.Default,
	}
	numericMatches := numericRe.FindAllStringSubmatch(c.DataType, -1)
	if len(numericMatches) > 0 {
		numericPrecision, err := strconv.Atoi(numericMatches[0][1])
		if err != nil {
			panic(err)
		}
		numericScale, err := strconv.Atoi(numericMatches[0][2])
		if err != nil {
			panic(err)
		}
		c.DataType = "numeric"
		c.NumericPrecision = numericPrecision
		c.NumericScale = numericScale
	}
	return db.ColumnNode{
		Column: c,
	}
}

func convertIndex(tableName, indexName string, v Index) db.IndexNode {
	return db.IndexNode{
		Index: db.Index{
			TableName: tableName,
			IndexName: indexName,
			Exprs:     v.On,
			Unique:    v.Unique,
		},
	}
}

func newPrimaryKey(tableName string, columnNames []string) db.IndexNode {
	return db.IndexNode{
		Index: db.Index{
			TableName: tableName,
			IndexName: fmt.Sprintf("%s_pkey", tableName),
			Exprs:     columnNames,
			Unique:    true,
			Primary:   true,
		},
	}
}
