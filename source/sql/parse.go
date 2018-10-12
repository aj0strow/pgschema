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
			break
		}
		if item.typ == itemError {
			err = fmt.Errorf("pgschema: sql lex error: %s", item.val)
			break
		}
		if item.typ == itemEOF {
			break
		}
		if item.typ == itemToken {
			switch item.val {
			case "EXTENSION":
				extNode := db.ExtensionNode{}
				err = parseExtension(&extNode, items)
				databaseNode.ExtensionNodes = append(databaseNode.ExtensionNodes, extNode)
			case "SCHEMA":
				schemaNode := db.SchemaNode{}
				err = parseSchema(&schemaNode.Schema, items)
				databaseNode.SchemaNodes = append(databaseNode.SchemaNodes, schemaNode)
			default:
				err = fmt.Errorf("pgschema: sql parse error: unexpected token")
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

func parseSchema(schema *db.Schema, items chan item) (err error) {
	schema.SchemaName, err = parseToken(items)
	if err != nil {
		return
	}
	err = parseColon(items)
	return
}

func parseToken(items chan item) (string, error) {
	item := <-items
	if item.typ == itemToken {
		return item.val, nil
	}
	return "", fmt.Errorf("pgschema: sql parse error: expected token")
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
