package pgschema

import (
	"github.com/jackc/pgx"
)

type PG interface {
	Query(string, ...interface{}) (*pgx.Rows, error)
	Exec(string, ...interface{}) (pgx.CommandTag, error)
}
