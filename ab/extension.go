package ab

import (
	"github.com/aj0strow/pgschema/db"
)

type ExtensionMatch struct {
	A *db.Extension
	B *db.Extension
}

func MatchExtensions(a, b []*db.Extension) []ExtensionMatch {
	var matches []ExtensionMatch
	fromA := map[string]bool{}
	for _, extA := range a {
		fromA[extA.ExtName] = true
		extB := findExtension(b, extA.ExtName)
		if extB != nil {
			matches = append(matches, ExtensionMatch{
				A: extA,
				B: extB,
			})
		} else {
			matches = append(matches, ExtensionMatch{
				A: extA,
				B: nil,
			})
		}
	}
	for _, extB := range b {
		if !fromA[extB.ExtName] {
			matches = append(matches, ExtensionMatch{
				A: nil,
				B: extB,
			})
		}
	}
	return matches
}

func findExtension(exts []*db.Extension, name string) *db.Extension {
	for _, ext := range exts {
		if ext.ExtName == name {
			return ext
		}
	}
	return nil
}
