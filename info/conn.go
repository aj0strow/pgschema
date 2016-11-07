package info

import (
	"github.com/jackc/pgx"
)

type Conn interface {
	Query(string, ...interface{}) (*pgx.Rows, error)
}
