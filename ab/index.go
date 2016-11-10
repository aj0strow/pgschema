package ab

import (
	"github.com/aj0strow/pgschema/db"
)

type IndexMatch struct {
	A *db.Index
	B *db.Index
}

func MatchIndexes(a, b []db.IndexNode) []IndexMatch {
	var matches []IndexMatch
	fromA := map[string]bool{}
	for _, nodeA := range a {
		indexA := nodeA.Index
		indexName := indexA.IndexName
		fromA[indexName] = true
		nodeB := findIndexNode(b, indexName)
		if nodeB == nil {
			matches = append(matches, IndexMatch{
				A: &indexA,
				B: nil,
			})
		} else {
			indexB := nodeB.Index
			matches = append(matches, IndexMatch{
				A: &indexA,
				B: &indexB,
			})
		}
	}
	for _, nodeB := range b {
		indexB := nodeB.Index
		if !fromA[indexB.IndexName] {
			matches = append(matches, IndexMatch{
				A: nil,
				B: &indexB,
			})
		}
	}
	return matches
}

func findIndexNode(nodes []db.IndexNode, name string) *db.IndexNode {
	for _, node := range nodes {
		if node.Index.IndexName == name {
			return &node
		}
	}
	return nil
}
