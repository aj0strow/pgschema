package psql

import (
	"fmt"

	"github.com/aj0strow/pgschema/db"
)

/*
pg_class (
	oid oid            -- object primary key
    relname text       -- provided name
);

pg_index (
	indexrelid oid     -- index pg_class foreign key
	indrelid oid       -- table pg_class foreign key
	indisprimary bool  -- it's actually a 't' or 'f' char
	indisunique bool   -- also a 't' or 'f' char
);

pg_catalog.pg_description (
	objoid oid         -- object foreign key
    description text   -- the comment
);
*/

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
			c.relname AS index_name
		FROM pg_index as ix
		JOIN pg_class as c ON (c.oid = ix.indexrelid)
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
		err := rows.Scan(&index.IndexName)
		if err != nil {
			return nil, err
		}
		index.TableName = tableName
		indexes = append(indexes, index)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return indexes, nil
}
