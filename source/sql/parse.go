package sql

import (
	"fmt"
	"github.com/aj0strow/pgschema/db"
)

func parse(input string) (db.DatabaseNode, error) {
	return parseSQL(lex(input))
}

func parseSQL(items chan item) (databaseNode db.DatabaseNode, err error) {
	for item := range items {
		if err != nil {
			return
		}
		if item.typ == itemError {
			err = fmt.Errorf("pgschema: sql lex error: %s", item.val)
			break
		}
		if item.typ == itemEOF {
			break
		}
		if item.typ == itemToken {
			if item.val == "EXTENSION" {
				extNode := db.ExtensionNode{}
				err = parseExtension(&extNode, items)
				databaseNode.ExtensionNodes = append(databaseNode.ExtensionNodes, extNode)
			}
		}
	}
	return
}

func parseExtension(extNode *db.ExtensionNode, items chan item) (err error) {
	extNode.Extension.ExtName, err = parseString(items)
	if err != nil {
		return
	}
	err = parseColon(items)
	return
}

func parseString(items chan item) (string, error) {
	item := <-items
	if item.typ == itemString {
		return item.val, nil
	}
	return "", fmt.Errorf("pgschema: sql parse error: expected string")
}

func parseColon(items chan item) error {
	item := <-items
	if item.typ == itemSpecial && item.val == ";" {
		return nil
	}
	return fmt.Errorf("pgschema: sql parse error: expected colon")
}
