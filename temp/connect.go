package temp

import (
	"fmt"
	"github.com/jackc/pgx"
)

const duplicateSchema = "42P06"
const duplicateKey = "23505"

type Conn struct {
	*pgx.Conn
	SchemaName string
}

func Connect(database string) (*Conn, error) {
	config := pgx.ConnConfig{
		Host:     "localhost",
		Database: database,
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		return nil, err
	}
	var schema string
	for {
		schema = randSchemaName()
		_, err = conn.Exec(fmt.Sprintf(`CREATE SCHEMA %s`, schema))
		if err == nil {
			break
		}
		if pgErr, ok := err.(pgx.PgError); ok {
			if pgErr.Code == duplicateSchema {
				continue
			}
			if pgErr.Code == duplicateKey {
				continue
			}
		}
		return nil, err
	}
	_, err = conn.Exec(fmt.Sprintf(`SET search_path TO %s`, schema))
	if err != nil {
		return nil, err
	}
	tmp := &Conn{
		Conn:       conn,
		SchemaName: schema,
	}
	return tmp, nil
}

func (tmp *Conn) Close() error {
	_, err := tmp.Exec(fmt.Sprintf(`DROP SCHEMA %s CASCADE`, tmp.SchemaName))
	if err != nil {
		return err
	}
	return tmp.Conn.Close()
}
