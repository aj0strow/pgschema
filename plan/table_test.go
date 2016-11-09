package plan

import (
	"github.com/aj0strow/pgschema/info"
	"github.com/aj0strow/pgschema/tree"
	"reflect"
	"testing"
)

func TestPlanTableMatch(t *testing.T) {
	type Test struct {
		Name       string
		TableMatch tree.TableMatch
		Changes    []Change
	}
	tests := []Test{
		Test{
			"create new table",
			tree.TableMatch{
				A: &info.Table{
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
			tree.TableMatch{
				A: nil,
				B: &info.Table{
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
