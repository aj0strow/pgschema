package sql

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	tt := []struct {
		input string
		items []item
	}{
		{
			"",
			[]item{
				{itemEOF, ""},
			},
		},
		{
			"-- user table",
			[]item{
				{itemEOF, ""},
			},
		},
		{
			"/* user table */",
			[]item{
				{itemEOF, ""},
			},
		},
		{
			"/* nested /* comments */ */",
			[]item{
				{itemEOF, ""},
			},
		},
		{
			";",
			[]item{
				{itemSpecial, ";"},
				{itemEOF, ""},
			},
		},
		{
			"table",
			[]item{
				{itemToken, "table"},
				{itemEOF, ""},
			},
		},
		{
			"SELECT * FROM MY_TABLE;",
			[]item{
				{itemToken, "SELECT"},
				{itemSpecial, "*"},
				{itemToken, "FROM"},
				{itemToken, "MY_TABLE"},
				{itemSpecial, ";"},
				{itemEOF, ""},
			},
		},
	}
	for _, tc := range tt {
		var items []item
		l := lex(tc.input)
		for item := range l.items {
			items = append(items, item)
		}
		if !reflect.DeepEqual(items, tc.items) {
			t.Errorf("%s\nWant:\n%v\nHave:\n%v\n", tc.input, tc.items, items)
		}
	}
}
