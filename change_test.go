package pgschema

import (
	"fmt"
	"github.com/jackc/pgx"
	"testing"
)

const SyntaxError = "42601"

func checkSyntax(pg PG, q string) error {
	_, err := pg.Exec(q)
	if pgErr, ok := err.(pgx.PgError); ok {
		if pgErr.Code != SyntaxError {
			return nil
		}
	}
	return err
}

func TestChanges(t *testing.T) {
	conn, err := Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	changes := []Change{
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
