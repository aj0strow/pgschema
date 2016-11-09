package plan

import (
	"github.com/aj0strow/pgschema/ab"
	"github.com/aj0strow/pgschema/db"
	"reflect"
	"testing"
)

func TestPlanTableMatch(t *testing.T) {
	type Test struct {
		Name       string
		TableMatch ab.TableMatch
		Changes    []Change
	}
	tests := []Test{
		Test{
			"create new table",
			ab.TableMatch{
				A: &db.Table{
					TableName: "users",
				},
				B: nil,
			},
			[]Change{
				CreateTable{"users"},
			},
		},
		Test{
			"drop old table",
			ab.TableMatch{
				A: nil,
				B: &db.Table{
					TableName: "customers",
				},
			},
			[]Change{
				DropTable{"customers"},
			},
		},
	}
	for _, test := range tests {
		changes := TableChanges(test.TableMatch)
		if !reflect.DeepEqual(changes, test.Changes) {
			t.Errorf("planTableMatch => %s", test.Name)
		}
	}
}
