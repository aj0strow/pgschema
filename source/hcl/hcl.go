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

func ParseBytes(bs []byte) (*db.Database, error) {
	var database Database
	err := hcl.Unmarshal(bs, &database)
	if err != nil {
		return nil, err
	}
	return convertDatabase(database), nil
}

func convertDatabase(v Database) *db.Database {
	var exts []*db.Extension
	for ek := range v.Extension {
		exts = append(exts, convertExtension(ek))
	}
	var schemas []*db.Schema
	for sk, sv := range v.Schema {
		schemas = append(schemas, convertSchema(sk, sv))
	}
	database := &db.Database{
		Extensions: exts,
		Schemas:    schemas,
	}
	return database
}

func convertExtension(k string) *db.Extension {
	return &db.Extension{
		ExtName: k,
	}
}

func convertSchema(k string, v Schema) *db.Schema {
	var tables []*db.Table
	for tableName, tv := range v.Table {
		tables = append(tables, convertTable(tableName, tv))
	}
	return &db.Schema{
		SchemaName: k,
		Tables:     tables,
	}
}

func convertTable(tableName string, v Table) *db.Table {
	var (
		columns []*db.Column
		indexes []*db.Index
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
	table := &db.Table{
		TableName: tableName,
		Columns:   columns,
		Indexes:   indexes,
	}
	return table
}

var numericRe = regexp.MustCompile(`numeric\((\d+),\s*(\d+)\)`)
var arrayRe = regexp.MustCompile(`^(.+)(?:\[\d*\])+$`)

func convertColumn(k string, v Column) *db.Column {
	c := &db.Column{
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
	arrayMatches := arrayRe.FindAllStringSubmatch(c.DataType, -1)
	if len(arrayMatches) > 0 {
		c.DataType = arrayMatches[0][1]
		c.Array = true
	}
	return c
}

func convertIndex(tableName, indexName string, v Index) *db.Index {
	return &db.Index{
		TableName: tableName,
		IndexName: indexName,
		Exprs:     v.On,
		Unique:    v.Unique,
	}
}

func newPrimaryKey(tableName string, columnNames []string) *db.Index {
	return &db.Index{
		TableName: tableName,
		IndexName: fmt.Sprintf("%s_pkey", tableName),
		Exprs:     columnNames,
		Unique:    true,
		Primary:   true,
	}
}
