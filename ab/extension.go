package ab

import (
	"github.com/aj0strow/pgschema/db"
)

type ExtensionMatch struct {
	A *db.Extension
	B *db.Extension
}

func MatchExtensions(a, b []db.ExtensionNode) []ExtensionMatch {
	var matches []ExtensionMatch
	fromA := map[string]bool{}
	for _, nodeA := range a {
		extA := nodeA.Extension
		fromA[extA.ExtName] = true
		nodeB := findExtensionNode(b, extA.ExtName)
		if nodeB != nil {
			extB := nodeB.Extension
			matches = append(matches, ExtensionMatch{
				A: &extA,
				B: &extB,
			})
		} else {
			matches = append(matches, ExtensionMatch{
				A: &extA,
				B: nil,
			})
		}
	}
	for _, nodeB := range b {
		extB := nodeB.Extension
		if !fromA[extB.ExtName] {
			matches = append(matches, ExtensionMatch{
				A: nil,
				B: &extB,
			})
		}
	}
	return matches
}

func findExtensionNode(nodes []db.ExtensionNode, name string) *db.ExtensionNode {
	for _, node := range nodes {
		if node.Extension.ExtName == name {
			return &node
		}
	}
	return nil
}
