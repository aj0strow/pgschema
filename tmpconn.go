package pgschema

import (
	"fmt"
	"github.com/jackc/pgx"
)

var (
	counter = 0
)

func randSchemaName() string {
	counter += 1
	return fmt.Sprintf("db%d", counter)
}

type TmpConn struct {
	*pgx.Conn
	Schema string
}

func Connect(database string) (*TmpConn, error) {
	config := pgx.ConnConfig{
		Host:     "localhost",
		Database: database,
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		return nil, err
	}
	schema := randSchemaName()
	_, err = conn.Exec(fmt.Sprintf(`CREATE SCHEMA %s`, schema))
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec(fmt.Sprintf(`SET search_path TO %s`, schema))
	if err != nil {
		return nil, err
	}
	tmp := &TmpConn{
		Conn:   conn,
		Schema: schema,
	}
	return tmp, nil
}

func (tmp *TmpConn) Close() error {
	_, err := tmp.Exec(fmt.Sprintf(`DROP SCHEMA %s CASCADE`, tmp.Schema))
	if err != nil {
		return err
	}
	return tmp.Conn.Close()
}
