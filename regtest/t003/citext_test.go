package t001

import (
	"github.com/aj0strow/pgschema/regtest"
	"github.com/aj0strow/pgschema/temp"
	"testing"
)

func TestCitext(t *testing.T) {
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
}
