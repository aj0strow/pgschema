package t001

import (
	"github.com/aj0strow/pgschema/regtest"
	"github.com/aj0strow/pgschema/temp"
	"testing"
)

func TestArray(t *testing.T) {
	conn, err := temp.Connect("pgschema")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	rt := regtest.RegressionTest{
		Conn:       conn,
		SetupFile:  "./setup.sql",
		SourceFile: "./schema.hcl",
	}
	have, err := rt.Run()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range have {
		t.Errorf(c.String())
	}
	tag, err := conn.Exec(`INSERT INTO finance_data (sent5m) VALUES ($1)`, []float64{1.5, 2.5, 3.5})
	if err != nil {
		t.Fatal(err)
	}
	if tag.RowsAffected() != 1 {
		t.Errorf("did not insert rows")
	}
}
