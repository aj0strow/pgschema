package sql

import (
	"fmt"
	"github.com/aj0strow/pgschema/db"
)

type parser struct {
	index int
	items []item
}

func (p *parser) Has() bool {
	return p.index < len(p.items)-1
}

func (p *parser) Next() item {
	i := p.items[p.index]
	if len(p.items) > p.index+1 {
		p.index += 1
	}
	return i
}

func (p *parser) Back() {
	if p.index > 0 {
		p.index -= 1
	}
}

func newParseErr(reason string) error {
	return fmt.Errorf("pgschema: sql parse error: %s", reason)
}

func parse(input string) (db.DatabaseNode, error) {
	p := newParser(lex(input))
	var databaseNode db.DatabaseNode
	var err = parseDatabaseNode(p, &databaseNode)
	return databaseNode, err
}

func newParser(items chan item) *parser {
	p := &parser{}
	for item := range items {
		p.items = append(p.items, item)
	}
	return p
}

func parseDatabaseNode(p *parser, databaseNode *db.DatabaseNode) error {
	var i item
	for p.Has() {
		i = p.Next()
		switch i.typ {
		case itemError:
			return newParseErr(i.val)
		case itemEOF:
			return nil
		case itemToken:
			switch i.val {
			case "EXTENSION":
				var extNode db.ExtensionNode
				if err := parseExtensionNode(p, &extNode); err != nil {
					return err
				}
				databaseNode.ExtensionNodes = append(databaseNode.ExtensionNodes, extNode)
			case "SCHEMA":
				var schemaNode db.SchemaNode
				if err := parseSchemaNode(p, &schemaNode); err != nil {
					return err
				}
				databaseNode.SchemaNodes = append(databaseNode.SchemaNodes, schemaNode)
			case "TABLE":
				if len(databaseNode.SchemaNodes) == 0 {
					return newParseErr("missing schema for table definition")
				}
				last := len(databaseNode.SchemaNodes) - 1
				schemaNode := databaseNode.SchemaNodes[last]
				tableNode := db.TableNode{}
				if err := parseTableNode(p, &tableNode); err != nil {
					return err
				}
				schemaNode.TableNodes = append(schemaNode.TableNodes, tableNode)
				databaseNode.SchemaNodes[last] = schemaNode
			default:
				return newParseErr("unexpected token `" + i.val + "`")
			}
		default:
			return newParseErr("expected token")
		}
	}
	return nil
}

func parseExtensionNode(p *parser, extNode *db.ExtensionNode) error {
	var i item
	i = p.Next()
	if i.typ == itemToken {
		extNode.Extension.ExtName = i.val
	} else {
		return newParseErr("expected extension name")
	}
	i = p.Next()
	if i.typ != itemSpecial || i.val != ";" {
		return newParseErr("expected semicolon")
	}
	return nil
}

func parseSchemaNode(p *parser, schemaNode *db.SchemaNode) error {
	var i item
	i = p.Next()
	if i.typ == itemToken {
		schemaNode.Schema.SchemaName = i.val
	} else {
		return newParseErr("expected schema name")
	}
	i = p.Next()
	if i.typ != itemSpecial || i.val != ";" {
		return newParseErr("expected semicolon")
	}
	return nil
}

func parseTableNode(p *parser, tableNode *db.TableNode) error {
	var i item
	i = p.Next()
	if i.typ == itemToken {
		tableNode.Table.TableName = i.val
	} else {
		return newParseErr("expected table name")
	}
	i = p.Next()
	if i.typ != itemSpecial || i.val != "(" {
		return newParseErr("expected open paren")
	}
	for p.Has() {
		i = p.Next()
		if i.typ == itemSpecial && i.val == ")" {
			break
		}
		p.Back()
		var columnNode db.ColumnNode
		if err := parseColumnNode(p, tableNode, &columnNode); err != nil {
			return err
		}
		tableNode.ColumnNodes = append(tableNode.ColumnNodes, columnNode)
	}
	i = p.Next()
	if i.typ != itemSpecial || i.val != ";" {
		return newParseErr("expected semicolon")
	}
	return nil
}

func parseColumnNode(p *parser, tableNode *db.TableNode, columnNode *db.ColumnNode) error {
	var i item
	i = p.Next()
	if i.typ == itemToken {
		columnNode.Column.ColumnName = i.val
	} else {
		return newParseErr("expected column name")
	}
	i = p.Next()
	if i.typ == itemToken {
		columnNode.Column.DataType = i.val
	} else {
		return newParseErr("expected column data type")
	}
	for p.Has() {
		i = p.Next()
		if i.typ == itemSpecial && i.val == ")" {
			p.Back()
			return nil
		}
		if i.typ == itemSpecial && i.val == "," {
			return nil
		}
		p.Back()
		if err := parseColumnConstraint(p, tableNode, columnNode); err != nil {
			return err
		}
	}
	return nil
}

func parseColumnConstraint(p *parser, tableNode *db.TableNode, columnNode *db.ColumnNode) error {
	var i item
	i = p.Next()
	if i.typ != itemToken {
		return newParseErr("expected column constraint")
	}
	if i.val == "NOT" {
		i = p.Next()
		if i.typ == itemToken && i.val == "NULL" {
			columnNode.Column.NotNull = true
			return nil
		} else {
			return newParseErr("expected keyword null")
		}
	}
	if i.val == "PRIMARY" {
		i = p.Next()
		if i.typ == itemToken && i.val == "KEY" {
			var indexNode db.IndexNode
			indexNode.Index = db.Index{
				TableName: tableNode.Table.TableName,
				IndexName: tableNode.Table.TableName + "_pkey",
				Exprs:     []string{columnNode.Column.ColumnName},
				Unique:    true,
				Primary:   true,
			}
			tableNode.IndexNodes = append(tableNode.IndexNodes, indexNode)
			return nil
		} else {
			return newParseErr("expected keyword key")
		}
	}
	if i.val == "UNIQUE" {
		var indexNode db.IndexNode
		indexNode.Index = db.Index{
			TableName: tableNode.Table.TableName,
			IndexName: tableNode.Table.TableName + "_" + columnNode.Column.ColumnName + "_key",
			Exprs:     []string{columnNode.Column.ColumnName},
			Unique:    true,
		}
		tableNode.IndexNodes = append(tableNode.IndexNodes, indexNode)
		return nil
	}
	if i.val == "DEFAULT" {
		return parseColumnDefault(p, columnNode)
	}
	return newParseErr("unexpected column constraint")
}

func parseColumnDefault(p *parser, columnNode *db.ColumnNode) error {
	var i item
	i = p.Next()
	if i.typ == itemString {
		columnNode.Column.Default = "'" + i.val + "'"
		return nil
	}
	if i.typ == itemNumber {
		columnNode.Column.Default = i.val
		return nil
	}
	if i.typ == itemToken {
		expr := i.val
		i = p.Next()
		if i.typ == itemSpecial && i.val == "(" {
			expr += i.val
			for p.Has() {
				i = p.Next()
				expr += i.val
				if i.typ == itemSpecial && i.val == ")" {
					break
				}
			}
		}
		columnNode.Column.Default = expr
		return nil
	}
	return newParseErr("expected default expression")
}
