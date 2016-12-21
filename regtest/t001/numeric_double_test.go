package t001

import (
	"github.com/aj0strow/pgschema/order"
	"github.com/aj0strow/pgschema/regtest"
	"github.com/aj0strow/pgschema/temp"
	"testing"
)

func TestNumericDouble(t *testing.T) {
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

	want := []order.Change{
		order.AlterTable{
			SchemaName: conn.SchemaName,
			TableName:  "accounts",
			Change: order.AlterColumn{
				ColumnName: "balance",
				Change: order.SetDataType{
					DataType: "double precision",
				},
			},
		},
	}

	for _, wc := range want {
		found := false
		for _, hc := range have {
			if wc.String() == hc.String() {
				found = true
			}
		}
		if !found {
			t.Errorf(wc.String())
		}
	}
}
