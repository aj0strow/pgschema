package order

import (
	"fmt"
	"testing"

	"github.com/aj0strow/pgschema/temp"
	"github.com/jackc/pgx"
)

const SyntaxError = "42601"

func checkSyntax(conn *temp.Conn, q string) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(q)
	if pgErr, ok := err.(pgx.PgError); ok {
		if pgErr.Code != SyntaxError {
			return nil
		}
	}
	return err
}

func TestSyntax(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	changes := []Change{
		CreateSchema{
			SchemaName: conn.SchemaName,
		},
		CreateTable{
			SchemaName: "public",
			TableName:  "users",
		},
		DropTable{
			SchemaName: "public",
			TableName:  "customers",
		},
		AlterTable{
			SchemaName: "public",
			TableName:  "laywers",
			Change: AddColumn{
				ColumnName: "name",
				DataType:   "text",
			},
		},
		AlterTable{
			SchemaName: "public",
			TableName:  "programmers",
			Change:     DropColumn{"weekends"},
		},
		AlterTable{
			SchemaName: "public",
			TableName:  "toronto",
			Change: AlterColumn{
				ColumnName: "house",
				Change: SetDataType{
					DataType: "text",
				},
			},
		},
		AlterTable{
			SchemaName: "public",
			TableName:  "address",
			Change: AlterColumn{
				ColumnName: "street",
				Change: SetDataType{
					DataType: "integer",
					Using:    "trim(street)::integer",
				},
			},
		},
		AlterTable{
			SchemaName: "public",
			TableName:  "address",
			Change: AlterColumn{
				ColumnName: "street",
				Change:     SetNotNull,
			},
		},
		AlterTable{
			SchemaName: "public",
			TableName:  "address",
			Change: AlterColumn{
				ColumnName: "street",
				Change:     DropNotNull,
			},
		},
		AlterTable{
			SchemaName: "public",
			TableName:  "address",
			Change: AlterColumn{
				ColumnName: "street",
				Change:     SetDefault{"'placeholder'"},
			},
		},
		AlterTable{
			SchemaName: "public",
			TableName:  "address",
			Change: AlterColumn{
				ColumnName: "street",
				Change:     DropDefault,
			},
		},
		CreateIndex{
			SchemaName: "public",
			TableName:  "users",
			IndexName:  "users_email_idx",
			Exprs:      []string{"lower(email)"},
		},
		CreateIndex{
			SchemaName: "public",
			TableName:  "users",
			IndexName:  "users_email_idx",
			Exprs:      []string{"lower(email)"},
			Unique:     true,
		},
		DropIndex{
			SchemaName: "public",
			IndexName:  "users_email_key",
		},
	}
	for _, change := range changes {
		err := checkSyntax(conn, change.String())
		if err != nil {
			fmt.Printf("%+v\n", change)
			fmt.Printf("%s\n", change)
			t.Fatal(err)
		}
	}
}
