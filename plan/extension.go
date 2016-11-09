package plan

import (
	"github.com/aj0strow/pgschema/ab"
)

var keepExtension = map[string]bool{
	"plpgsql": true,
}

func ExtensionChanges(extMatch ab.ExtensionMatch) []Change {
	var cs []Change
	a, b := extMatch.A, extMatch.B
	if a == nil {
		if !keepExtension[b.ExtName] {
			return append(cs, DropExtension{ExtName: b.ExtName})
		}
		return cs
	}
	if b == nil {
		return append(cs, CreateExtension{ExtName: a.ExtName})
	}
	return cs
}
