package info

import (
	"github.com/aj0strow/pgschema/db"
)

func LoadExtensions(conn Conn) ([]db.Extension, error) {
	rows, err := conn.Query(`
		SELECT extname
		FROM pg_extension
		ORDER BY extname ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var exts []db.Extension
	for rows.Next() {
		ext := db.Extension{}
		rows.Scan(&ext.ExtName)
		exts = append(exts, ext)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return exts, nil
}
