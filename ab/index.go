package ab

import (
	"github.com/aj0strow/pgschema/db"
)

type IndexMatch struct {
	A *db.Index
	B *db.Index
}

func MatchIndexes(a, b []*db.Index) []IndexMatch {
	var matches []IndexMatch
	fromA := map[string]bool{}
	for _, indexA := range a {
		indexName := indexA.IndexName
		fromA[indexName] = true
		indexB := findIndex(b, indexName)
		if indexB == nil {
			matches = append(matches, IndexMatch{
				A: indexA,
				B: nil,
			})
		} else {
			matches = append(matches, IndexMatch{
				A: indexA,
				B: indexB,
			})
		}
	}
	for _, indexB := range b {
		if !fromA[indexB.IndexName] {
			matches = append(matches, IndexMatch{
				A: nil,
				B: indexB,
			})
		}
	}
	return matches
}

func findIndex(idxs []*db.Index, name string) *db.Index {
	for _, idx := range idxs {
		if idx.IndexName == name {
			return idx
		}
	}
	return nil
}
