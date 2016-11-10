package plan

import (
	"fmt"
	"github.com/aj0strow/pgschema/temp"
	"github.com/jackc/pgx"
	"testing"
)

const SyntaxError = "42601"

func checkSyntax(conn *temp.Conn, q string) error {
	_, err := conn.Exec(q)
	if pgErr, ok := err.(pgx.PgError); ok {
		if pgErr.Code != SyntaxError {
			return nil
		}
	}
	return err
}

func TestChanges(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	changes := []Change{
		CreateSchema{conn.SchemaName},
		CreateTable{"users"},
		DropTable{"customers"},
		AlterTable{
			"laywers",
			AddColumn{"name", "text"},
		},
		AlterTable{
			"programmers",
			DropColumn{"weekends"},
		},
		AlterTable{
			"toronto",
			AlterColumn{
				"house",
				SetDataType{"text"},
			},
		},
		AlterTable{
			"address",
			AlterColumn{
				"street",
				CastDataType{
					Using:       "trim(street)::integer",
					SetDataType: SetDataType{"integer"},
				},
			},
		},
		AlterTable{
			"address",
			AlterColumn{
				"street",
				SetNotNull{},
			},
		},
		AlterTable{
			"address",
			AlterColumn{
				"street",
				DropNotNull{},
			},
		},
		DropIndex{
			"users_pkey",
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
