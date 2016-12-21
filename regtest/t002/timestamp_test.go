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

	negate := []order.Change{
		order.AlterTable{
			SchemaName: conn.SchemaName,
			TableName:  "users",
			Change: order.AlterColumn{
				ColumnName: "created_at",
				Change: order.SetDataType{
					DataType: "timestamp",
				},
			},
		},
		order.AlterTable{
			SchemaName: conn.SchemaName,
			TableName:  "customers",
			Change: order.AlterColumn{
				ColumnName: "created_at",
				Change: order.SetDataType{
					DataType: "timestamptz",
				},
			},
		},
	}

	for _, nc := range negate {
		found := false
		for _, hc := range have {
			if nc.String() == hc.String() {
				found = true
			}
		}
		if found {
			t.Errorf("invalid change")
			t.Errorf(nc.String())
		}
	}
}
