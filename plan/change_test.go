package plan

import (
	"fmt"
	"github.com/aj0strow/pgschema/temp"
	"github.com/jackc/pgx"
	"testing"
)

type Conn interface {
	Exec(string, ...interface{}) (pgx.CommandTag, error)
}

const SyntaxError = "42601"

func checkSyntax(db Conn, q string) error {
	_, err := db.Exec(q)
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
		CreateSchema{"v1"},
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
