// The psql package queries postgres meta tables to create a database db.
package psql

import (
	"github.com/aj0strow/pgschema/db"
)

func LoadDatabase(conn Conn) (*db.Database, error) {
	extensions, err := LoadExtensions(conn)
	if err != nil {
		return nil, err
	}
	schemas, err := LoadSchemas(conn)
	if err != nil {
		return nil, err
	}
	database := &db.Database{
		Extensions: extensions,
		Schemas:    schemas,
	}
	return database, nil
}
