package pgschema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlanTables(t *testing.T) {
	type Test struct {
		Name    string
		A, B    []Table
		Changes []Change
	}
	tests := []Test{
		Test{
			"no tables",
			[]Table{},
			[]Table{},
			[]Change{},
		},
		Test{
			"no changes",
			[]Table{
				Table{
					TableName: "customers",
				},
			},
			[]Table{
				Table{
					TableName: "customers",
				},
			},
			[]Change{},
		},
		Test{
			"create one table",
			[]Table{
				Table{
					TableName: "users",
				},
			},
			[]Table{},
			[]Change{
				CreateTable{
					TableName: "users",
				},
			},
		},
		Test{
			"drop one table",
			[]Table{},
			[]Table{
				Table{
					TableName: "events",
				},
			},
			[]Change{
				DropTable{
					TableName: "events",
				},
			},
		},
	}
	for _, test := range tests {
		changes := PlanTables(test.A, test.B)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("%s\n", test.Name)
			fmt.Printf("want: %+v\n", test.Changes)
			fmt.Printf("have: %+v\n", changes)
		}
	}
}
