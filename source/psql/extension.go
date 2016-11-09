package psql

import (
	"github.com/aj0strow/pgschema/db"
)

func LoadExtensionNodes(conn Conn) ([]db.ExtensionNode, error) {
	exts, err := LoadExtensions(conn)
	if err != nil {
		return nil, err
	}
	nodes := make([]db.ExtensionNode, len(exts))
	for i := range exts {
		nodes[i] = db.ExtensionNode{
			Extension: exts[i],
		}
	}
	return nodes, nil
}

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
