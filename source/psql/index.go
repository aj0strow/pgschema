package psql

import (
	"fmt"

	"github.com/aj0strow/pgschema/db"
)

func LoadIndexNodes(conn Conn, schema db.Schema, table db.Table) ([]db.IndexNode, error) {
	indexes, err := LoadIndexes(conn, schema.SchemaName, table.TableName)
	if err != nil {
		return nil, err
	}
	indexNodes := make([]db.IndexNode, len(indexes))
	for i := range indexes {
		indexNodes[i] = db.IndexNode{
			Index: indexes[i],
		}
	}
	return indexNodes, nil
}

func LoadIndexes(conn Conn, schemaName, tableName string) ([]db.Index, error) {
	q := fmt.Sprintf(`
		SELECT
			c.relname,
			ix.indisunique,
			ix.indisprimary
		FROM pg_index AS ix
		JOIN pg_class AS c ON (c.oid = ix.indexrelid)
		WHERE ix.indrelid = '%s.%s'::regclass
	`, schemaName, tableName)
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var indexes []db.Index
	for rows.Next() {
		var index db.Index
		err := rows.Scan(&index.IndexName, &index.Unique, &index.Primary)
		if err != nil {
			return nil, err
		}
		index.TableName = tableName
		indexes = append(indexes, index)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for i := range indexes {
		index := indexes[i]
		if index.Primary {
			names, err := LoadIndexExprs(conn, schemaName, index.IndexName)
			if err != nil {
				return nil, err
			}
			index.Exprs = names
			indexes[i] = index
		}
	}
	return indexes, nil
}

func LoadIndexExprs(conn Conn, schemaName, indexName string) ([]string, error) {
	q := fmt.Sprintf(`
		SELECT at.attname
		FROM pg_index AS ix
		JOIN pg_attribute AS at
		  ON at.attrelid = ix.indrelid AND at.attnum = ANY(ix.indkey)
		WHERE ix.indexrelid = '%s.%s'::regclass
		  AND ix.indisprimary
	`, schemaName, indexName)
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return names, nil
}
